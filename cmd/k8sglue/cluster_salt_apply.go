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

	"github.com/spf13/cobra"
)

// clusterSaltApplyCmd represents the salt command
var clusterSaltApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("salt called")
		return errCommandNotImplemented
	},
}

func init() {
	clusterCmd.AddCommand(clusterSaltApplyCmd)
	clusterSaltApplyCmd.Flags().StringVar(&saltStatesDir, "salt-states", "./salt", "Path to the `salt/` directory which contains the salt states to be copied to each salt-master(s)")
}
