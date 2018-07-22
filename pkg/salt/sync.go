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
	"os"
	"path"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/executor"
	"github.com/galexrt/k8sglue/pkg/util"
)

// Sync syncs current local `salt/` directory to the salt-master(s).
func Sync(masters []string) error {
	if err := util.Symlink(path.Join(config.Cfg.SaltDir), path.Join(config.Cfg.TempDir, "data", "k8sglue-salt")); err != nil {
		return err
	}
	defer os.Remove(path.Join(config.Cfg.TempDir, "data", "k8sglue-salt"))

	args := append(getSaltSSHDefaultArgs(),
		generateTargetFlags(masters)...,
	)
	args = append(args,
		"state.single",
		"file.recurse",
		"name=/srv",
		"source=salt://k8sglue-salt",
		"dir_mode=0750",
		"user=root",
		"group=root",
	)

	return executor.ExecOutToLog("salt-ssh copy salt", SaltSSHCommand, args)
}
