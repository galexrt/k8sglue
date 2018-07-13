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

	"github.com/galexrt/k8sglue/pkg/cmd/machines"
	"github.com/galexrt/k8sglue/pkg/cmd/salt"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/models"
	"github.com/galexrt/k8sglue/pkg/util"
	"github.com/spf13/cobra"
)

// clusterSaltCertsCmd represents the init command
var clusterSaltCertsCmd = &cobra.Command{
	Use:   "certs",
	Short: "Create certs for salt-master(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("certs called")
		config.Cfg.TempDir = "."

		masters := config.Cfg.Cluster.MachineList.GetMachinesByRole(models.Roles{
			Salt: &models.RolesSalt{
				Master: &util.True,
			},
		})
		logger.Debugf("master machines found by `salt.master: true` role: %+v\n", masters)
		if len(masters) == 0 {
			return fmt.Errorf("no machines with `roles.salt.master: true` found")
		}
		certPath, keyPath, err := salt.Certs(machines.GetAddressesFromMachines(masters))
		logger.Infof("cert: %s", certPath)
		logger.Infof("key: %s", keyPath)

		return err
	},
}

func init() {
	clusterSaltCmd.AddCommand(clusterSaltCertsCmd)
}
