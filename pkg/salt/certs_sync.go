/*
Copyright (c) 2018 Alexander Trost <galexrt@googlemail.com>. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package salt

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/executor"
)

// CertsSync sync cert and key to the salt-master(s)
func CertsSync(masters []string, cert, key []byte) error {
	for name, content := range map[string][]byte{
		"salt-master.crt": cert,
		"salt-master.key": key,
	} {
		if err := ioutil.WriteFile(path.Join(config.Cfg.TempDir, "data", name), content, 0600); err != nil {
			return err
		}
	}

	args := append(getSaltSSHDefaultArgs(),
		generateTargetFlags(masters)...,
	)
	args = append(args,
		"state.single",
		"file.managed",
		"makedirs=True",
		"mode=0600",
		"user=root",
		"group=root",
		"subdir=True",
		"hide_output=True",
	)

	for _, file := range []string{"salt-master.crt", "salt-master.key"} {
		cmdArgs := append(args,
			fmt.Sprintf("name=%s", path.Join("/etc/salt/ssl", file)),
			fmt.Sprintf("source=salt://%s", file),
		)
		if err := executor.ExecOutToLog("salt-ssh copy certs", SaltSSHCommand, cmdArgs); err != nil {
			return err
		}
		if err := os.Remove(path.Join(config.Cfg.TempDir, "data", file)); err != nil {
			return err
		}
	}

	return nil
}
