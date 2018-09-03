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
	"github.com/galexrt/k8sglue/pkg/executor"
)

// HighState definition of high state for `state.apply` call here
const HighState = ""

// SSHApply trigger salt-ssh highstate using salt-ssh on the salt-master(s)
func SSHApply(machines []string, slsFiles string) error {
	args := append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh", "state.apply")
	if slsFiles != "" {
		args = append(args, slsFiles)
	}

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

	return executor.ExecOutToLog("salt-ssh copy ssh key", SaltSSHCommand, args)
}
