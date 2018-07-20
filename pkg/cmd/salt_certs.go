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

	"github.com/galexrt/k8sglue/pkg/salt"
	"github.com/spf13/cobra"
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

		return salt.Certs()
	},
}

func init() {
	saltCmd.AddCommand(saltCertsCmd)
}
