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
	"os"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/spf13/cobra"
)

var logLevelRaw string
var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/cmd/k8sglue", "root")

var errCommandNotImplemented = fmt.Errorf("command or subcommand has not been implemented yet")

var banner = ` _   ___          _
| |_( _ )___ __ _| |_  _ ___
| / / _ (_-</ _` + "`" + ` | | || / -_)
|_\_\___/__/\__, |_|\_,_\___|
             |___/
Made by Alexander Trost
=======================
`

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8sglue",
	Short: `"Kleber" for Kubernetes in public and private (cloud) environments with Salt, kubeadm and some magic glue.`,
	Long: `k8sglue is a project which utilizes Kubernetes kubeadm to do the Kubernetes
installation/provisioning with Saltstack.

But with some magic allowing for simple integration into any server deployment
that in the end is able to spit out a list of machines to use.

For more information refer to the README.md and the docs.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Init(cmd.Name())
		SetLogLevel()
		fmt.Print(banner)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevelRaw, "log-level", "INFO", "Set log level")
}

func main() {
	Execute()
}

// SetLogLevel parses the raw log level and sets it as the global log level
func SetLogLevel() {
	// parse given log level string then set up corresponding global logging level
	ll, err := capnslog.ParseLevel(logLevelRaw)
	if err != nil {
		logger.Warningf("failed to set log level %s. %+v", logLevelRaw, err)
	}
	fmt.Printf("TEST: %+v\n", logLevelRaw)
	config.Cfg.LogLevel = ll
	capnslog.SetGlobalLogLevel(config.Cfg.LogLevel)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}
