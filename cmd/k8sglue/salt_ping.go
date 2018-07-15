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

// saltPingCmd represents the ping command
var saltPingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("ping called")
		return salt.Ping()
	},
}

func init() {
	saltCmd.AddCommand(saltPingCmd)
}
