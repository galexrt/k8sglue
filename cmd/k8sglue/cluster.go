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
	"io/ioutil"
	"os"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath string

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// TODO Remove the conditional flag check when https://github.com/spf13/cobra/issues/655 has been resolved
		if viper.GetString("cluster") != "" {
			tempDir, err := ioutil.TempDir(os.TempDir(), cmd.Root().Name())
			if err != nil {
				return nil
			}

			return config.Load(configPath, tempDir)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clusterCmd)
	clusterCmd.PersistentFlags().StringVar(&configPath, "cluster", "", "Cluster config directory")
	clusterCmd.MarkPersistentFlagRequired("cluster")
	viper.BindPFlag("cluster", clusterCmd.PersistentFlags().Lookup("cluster"))
}
