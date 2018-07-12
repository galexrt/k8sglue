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

package main

import (
	"fmt"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/models"
	"github.com/galexrt/k8sglue/pkg/util"
	"github.com/spf13/cobra"
)

// clusterInitCmd represents the init command
var clusterInitCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("init called")

		masters := config.Cfg.MachineList.GetMachinesByRole(models.Roles{
			Salt: &models.RolesSalt{
				Master: &util.True,
			},
		})
		logger.Debugf("master machines found by `salt.master: true` role: %+v\n", masters)

		// TODO Check if salt-master(s) certs, already exist locally
		// TODO Generate salt-master(s) certificate
		/*
			saltCert, saltKey, err := cert.Generate([]string{"127.0.0.1"}, "", 24*time.Hour, false, 4096, "P521")
			if err != nil {
				return err
			}
		*/

		// TODO scp the salt-master certificate to the salt-master(s)

		// TODO scp the `salt/` directory to the salt-master(s)

		return nil
	},
}

func init() {
	clusterCmd.AddCommand(clusterInitCmd)
}
