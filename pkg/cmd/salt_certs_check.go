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

// saltCertsSyncmd represents the generate command
var saltCertsSyncmd = &cobra.Command{
	Use:   "check",
	Short: "Check the certificates on the salt-master(s).",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("salt certs sync called")
		if err := bootstrapCommand(cmd, true); err != nil {
			return err
		}

		masters := config.Cfg.Cluster.Salt.Roster.GetEntriesByRole("salt-master").GetNames()
		if len(masters) == 0 {
			return fmt.Errorf("no nodes with role salt-master found")
		}

		// TODO Load cert and key from file
		err := salt.CertsSync(masters, []byte{}, []byte{})
		if err != nil {
			return err
		}

		return errCommandNotImplemented
	},
}

func init() {
	saltCertsCmd.AddCommand(saltCertsSyncmd)

	saltCertsSyncmd.Flags().String("cert", "", "Path to certificate to sync")
	saltCertsSyncmd.Flags().String("key", "", "Path to key to sync")
	saltCertsSyncmd.MarkFlagRequired("cert")
	saltCertsSyncmd.MarkFlagRequired("key")

	viper.BindPFlag("cert", saltCertsSyncmd.Flags().Lookup("cert"))
	viper.BindPFlag("key", saltCertsSyncmd.Flags().Lookup("key"))
}
