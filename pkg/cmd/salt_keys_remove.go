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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// saltKeysRemoveCmd represents the remove command
var saltKeysRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("remove called")

		return nil
	},
}

func init() {
	saltKeysCmd.AddCommand(saltKeysRemoveCmd)
	saltKeysRemoveCmd.Flags().StringSlice("hosts", []string{}, "a list of hosts")
	saltKeysRemoveCmd.MarkFlagRequired("hosts")
	viper.BindPFlag("hosts", saltKeysRemoveCmd.Flags().Lookup("hosts"))
}
