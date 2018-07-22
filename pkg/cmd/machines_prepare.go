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

	"github.com/spf13/cobra"
)

// machinesPrepareCmd represents the prepare command
var machinesPrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare one or more nodes by using salt-ssh.",
	Long: `Prepare one or more nodes by using salt-ssh to run the "base" states ("common" and "salt-minion").
In the end the node's salt-minion must be connected to the salt-master(s).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("machines prepare called")
		if err := bootstrapCommand(cmd, true); err != nil {
			return err
		}

		// TODO run salt-ssh over state file `base`

		return errCommandNotImplemented
	},
}

func init() {
	machinesCmd.AddCommand(machinesPrepareCmd)
}
