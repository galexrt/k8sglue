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

package cmd

import (
	"fmt"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/salt"
	"github.com/spf13/cobra"
)

// saltSyncCmd represents the sync command
var saltSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync current (given) `salt` directory to all salt-master(s).",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("salt sync called")
		if err := bootstrapCommand(cmd, true); err != nil {
			return err
		}

		masters := config.Cfg.Machines.GetEntriesByRole("salt-master").GetHosts()
		if len(masters) == 0 {
			return fmt.Errorf("no nodes with role salt-master found")
		}
		return salt.Sync(masters)
	},
}

func init() {
	saltCmd.AddCommand(saltSyncCmd)
}
