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
	"time"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/salt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// saltCertsCmd represents the certs command
var saltCertsCmd = &cobra.Command{
	Use:   "certs",
	Short: "Generate and sync certs for the salt-master(s) (if needed and force-able by flag).",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("salt certs called")
		if err := bootstrapCommand(cmd, true); err != nil {
			return err
		}

		saltMasters := config.Cfg.Cluster.Salt.Roster.GetEntriesByRole("salt-master")
		masters := saltMasters.GetNames()
		if len(masters) == 0 {
			return fmt.Errorf("no nodes with role salt-master found")
		}

		if err := salt.Ping(masters); err != nil {
			return err
		}

		// Now that we know that the nodes with the salt-master role are reachable add the `Host` field of them
		names := masters
		names = append(names, saltMasters.GetHosts()...)

		if err := salt.Certs(masters, names, viper.GetDuration("valid-time"), viper.GetDuration("renew-time"), viper.GetBool("force")); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	saltCmd.AddCommand(saltCertsCmd)

	// 4 Years
	saltCertsCmd.Flags().Duration("valid-time", 35040*time.Hour, "Validity time for generated certificates")
	// 0.5 Years
	saltCertsCmd.Flags().Duration("renew-time", 4380*time.Hour, "Renew certificate if equal or less than time left on validity of generated certificates (not in use right now)")
	viper.BindPFlag("valid-time", saltCertsCmd.Flags().Lookup("valid-time"))
	viper.BindPFlag("renew-time", saltCertsCmd.Flags().Lookup("renew-time"))
}
