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
	"github.com/spf13/viper"
)

// saltApplyCmd represents the apply command
var saltApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Trigger salt (high) state.",
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("hosts", cmd.Flags().Lookup("hosts"))
		viper.BindPFlag("all", cmd.Flags().Lookup("all"))
		viper.BindPFlag("sls-files", cmd.Flags().Lookup("sls-files"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("salt apply called")
		if err := bootstrapCommand(cmd, true); err != nil {
			return err
		}

		hosts := viper.GetStringSlice("hosts")
		if len(hosts) == 0 && !viper.GetBool("all") {
			return fmt.Errorf("no all or host flag given")
		} else if viper.GetBool("all") {
			hosts = config.Cfg.Machines.GetHosts()
		}
		return salt.Apply(hosts, viper.GetString("sls-files"))
	},
}

func init() {
	saltCmd.AddCommand(saltApplyCmd)

	saltApplyCmd.Flags().StringSlice("hosts", []string{}, "a list of hosts comma separated")
	saltApplyCmd.Flags().Bool("all", false, "if all hosts in the cluster machines list should be used")
	saltApplyCmd.Flags().StringP("sls-files", "s", "", "Which SLS files to call for state.apply, if none given high state")
}
