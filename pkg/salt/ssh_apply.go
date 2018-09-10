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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/executor"
)

// HighState definition of high state for `state.apply` call here
const HighState = ""

// SSHApply trigger salt-ssh highstate using salt-ssh on the salt-master(s)
func SSHApply(machines []string, slsFiles string) error {
	args := append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines[0:0])...,
	)
	args = append(args, "--state-verbose=false", "--refresh", "state.apply")
	if err := executor.ExecOutToLog("salt-ssh state.apply first master", SaltSSHCommand, args); err != nil {
		return err
	}

	// TODO Copy `/etc/salt/pki/master/master_sign.*` to the other masters
	files := []string{
		"/etc/salt/pki/master/master_sign.pem",
		"/etc/salt/pki/master/master_sign.pub",
	}
	for _, file := range files {
		args = append(getSaltSSHDefaultArgs(),
			generateTargetFlags(machines)...,
		)
		args = append(args, "--state-verbose=false", "--refresh", "--static", "--out=json", "cp.get_file_str", file)
		out, err := executor.ExecStdoutByte(SaltSSHCommand, args)
		if err != nil {
			return err
		}
		var parsed map[string]string
		if err = json.Unmarshal(out, &parsed); err != nil {
			return err
		}

		if err = ioutil.WriteFile(
			path.Join(config.Cfg.TempDir, "data", filepath.Base(file)),
			[]byte(parsed[machines[0]]),
			0600,
		); err != nil {
			return err
		}
	}

	for _, file := range files {
		args = append(getSaltSSHDefaultArgs(),
			generateTargetFlags(machines)...,
		)
		args = append(args, "--state-verbose=false", "--refresh",
			"state.single",
			"file.managed",
			fmt.Sprintf("name=%s", strings.Replace(file, "master", "minion", 1)),
			fmt.Sprintf("source=salt://%s", filepath.Base(file)),
			"dir_mode=0700",
			"mode=0600",
			"user=root",
			"group=root",
		)
		if err := executor.ExecOutToLog("salt-ssh copy master sign files", SaltSSHCommand, args); err != nil {
			return err
		}
	}

	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh", "state.apply")
	if err := executor.ExecOutToLog("salt-ssh state.apply", SaltSSHCommand, args); err != nil {
		return err
	}

	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh",
		"state.single",
		"file.managed",
		"name=/etc/salt/ssh/id_rsa",
		"source=salt://ssh_id_rsa",
		"dir_mode=0600",
		"mode=600",
		"user=root",
		"group=root",
	)

	if err := executor.ExecOutToLog("salt-ssh copy ssh key", SaltSSHCommand, args); err != nil {
		return err
	}

	if err := Sync(machines); err != nil {
		return err
	}

	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh",
		"state.single",
		"cmd.run",
		"name=systemctl restart salt-master salt-minion",
	)

	return executor.ExecOutToLog("salt-ssh restart salt-minion", SaltSSHCommand, args)
}
