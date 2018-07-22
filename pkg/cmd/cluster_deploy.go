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
	"github.com/spf13/viper"
)

// clusterDeployCmd represents the deploy command
var clusterDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Run the orchestrated Kubernetes cluster state.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cluster deploy called")
		if err := bootstrapCommand(cmd, true); err != nil {
			return err
		}

		// TODO Use salt-ssh to run `salt-run state.orchestrate orch.kubernetes` on one of the salt-master(s)
		return errCommandNotImplemented
	},
}

func init() {
	clusterCmd.AddCommand(clusterDeployCmd)

	clusterDeployCmd.Flags().Bool("machines-prepare", false, "If the machines should be prepared first or just be used.")
	viper.BindPFlag("machines-prepare", clusterDeployCmd.Flags().Lookup("machines-prepare"))
}
