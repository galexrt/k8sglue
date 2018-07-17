// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/galexrt/k8sglue/pkg/cmd/salt"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/spf13/cobra"
)

// saltAcceptCmd represents the accept command
var saltAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("accept called")

		for hostname := range config.Cfg.Cluster.Salt.Masters {
			return salt.KeyAccept(hostname)
		}
		return nil
	},
}

func init() {
	saltKeysCmd.AddCommand(saltAcceptCmd)
}
