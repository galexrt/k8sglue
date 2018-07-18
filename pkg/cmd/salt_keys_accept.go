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

package cmd

import (
	"fmt"

	"github.com/galexrt/k8sglue/pkg/salt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// saltKeysAcceptCmd represents the accept command
var saltKeysAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("accept called")

		return salt.KeyAcceptList(viper.GetStringSlice("host"))
	},
}

func init() {
	saltKeysCmd.AddCommand(saltKeysAcceptCmd)
	saltKeysAcceptCmd.Flags().StringSlice("hosts", []string{}, "a list of hosts")
	saltKeysAcceptCmd.MarkFlagRequired("hosts")
	viper.BindPFlag("hosts", saltKeysAcceptCmd.Flags().Lookup("hosts"))
}
