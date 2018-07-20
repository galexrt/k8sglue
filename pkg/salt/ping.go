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

import "github.com/galexrt/k8sglue/pkg/executor"

// Ping runs `salt-ssh` `test.ping` on all salt-master(s).
func Ping(machines []string) error {
	args := append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "test.ping")

	return executor.ExecOutToLog("salt-ssh test.ping", SaltSSHCommand, args)
}
