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

	"github.com/galexrt/k8sglue/pkg/cmd/salt"
	"github.com/spf13/cobra"
)

// saltSecretsCmd represents the secrets command
var saltSecretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("salt certs called")
		cert, key, err := salt.Secrets()
		fmt.Printf("Certificate:\n%s\n", cert)
		fmt.Printf("Key:\n%s\n", key)
		return err
	},
}

func init() {
	saltCmd.AddCommand(saltSecretsCmd)
}
