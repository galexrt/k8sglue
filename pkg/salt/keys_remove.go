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
	"strings"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/executor"
)

var saltKeyRemoveMagic = `
machines=(
%s
)
for (( i = 0; i < ${#machines[@]}; ++i)); do
	salt-key --out=json -q -y -d "${machines[i]}"
	echo "Removed key for minion ${machines[i]}"
done`

// KeyRemove removes minions keys from salt-master(s)
func KeyRemove(machines []string) error {
	masters := config.Cfg.Machines.GetEntriesByRole("salt-master").GetNames()
	if len(masters) == 0 {
		return fmt.Errorf("no nodes with role salt-master found")
	}

	args := append(getSaltSSHDefaultArgs(),
		"-L",
		strings.Join(masters, ","),
		"cmd.run",
		fmt.Sprintf(saltKeyRemoveMagic, strings.Join(machines, "\n")),
	)

	return executor.ExecOutToLog("salt-ssh salt-key", SaltSSHCommand, args)
}
