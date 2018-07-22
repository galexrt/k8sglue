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

// saltCertsSyncCmd represents the sync command
var saltCertsSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync certificates for salt-master(s).",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("salt certs generate called")
		if err := bootstrapCommand(cmd, true); err != nil {
			return err
		}

		names := config.Cfg.Machines.GetEntriesByRole("salt-master").GetHosts()
		cert, key, err := salt.CertsGenerate(names, viper.GetDuration("valid-time"))
		if err != nil {
			return err
		}

		fmt.Printf("Certificate:\n%s\n", cert)
		fmt.Printf("Key:\n%s\n", key)

		return nil
	},
}

func init() {
	saltCertsCmd.AddCommand(saltCertsSyncCmd)

	saltCertsSyncCmd.Flags().String("cert", "", "Path to certificate to sync")
	saltCertsSyncCmd.Flags().String("key", "", "Path to key to sync")
	saltCertsSyncCmd.MarkFlagRequired("cert")
	saltCertsSyncCmd.MarkFlagRequired("key")

	viper.BindPFlag("cert", saltCertsSyncCmd.Flags().Lookup("cert"))
	viper.BindPFlag("key", saltCertsSyncCmd.Flags().Lookup("key"))
}
