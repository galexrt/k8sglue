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

// Package cmd contains the k8sglue commands with their subcommands.
package cmd

import (
	"fmt"
	"os"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/salt"
	"github.com/galexrt/k8sglue/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logLevelRaw string
var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/cmd/k8sglue", "root")

var errCommandNotImplemented = fmt.Errorf("command or subcommand has not been implemented yet")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8sglue",
	Short: `"Kleber" for Kubernetes in public and private (cloud) environments with Salt, kubeadm and some magic glue.`,
	Long: `k8sglue is a project which utilizes Kubernetes kubeadm to do the Kubernetes
installation/provisioning with Saltstack.

But with some magic allowing for simple integration into any server deployment
that in the end is able to spit out a list of machines to use.

For more information refer to the README.md and the docs.`,
}

func init() {
	viper.SetEnvPrefix("k8sglue")
	viper.AutomaticEnv()

	rootCmd.PersistentFlags().StringVar(&logLevelRaw, "log-level", "INFO", "Set log level")

	rootCmd.PersistentFlags().String("cluster", "", "Cluster settings file")
	rootCmd.PersistentFlags().String("machines", "", "Cluster machines file")
	rootCmd.PersistentFlags().String("force", "", "Force the actions being run")
	rootCmd.PersistentFlags().String("temp-dir", "/tmp/k8sglue", "Temp directory for temporary files")
	rootCmd.PersistentFlags().String("salt-dir", "salt/", "Path to the salt files containing directory (will be appended to the current work dir by default")

	viper.BindPFlag("cluster", rootCmd.PersistentFlags().Lookup("cluster"))
	viper.BindPFlag("machines", rootCmd.PersistentFlags().Lookup("machines"))
	viper.BindPFlag("force", rootCmd.PersistentFlags().Lookup("force"))
	viper.BindPFlag("temp-dir", rootCmd.PersistentFlags().Lookup("temp-dir"))
	viper.BindPFlag("salt-dir", rootCmd.PersistentFlags().Lookup("salt-dir"))
}

// SetLogLevel parses the raw log level and sets it as the global log level
func SetLogLevel() {
	// parse given log level string then set up corresponding global logging level
	ll, err := capnslog.ParseLevel(logLevelRaw)
	if err != nil {
		logger.Warningf("failed to set log level %s. %+v", logLevelRaw, err)
	}
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

func bootstrapCommand(cmd *cobra.Command, prepareSaltSSH bool) error {
	config.Init("k8sglue")
	if cmd.Name() != "help" {
		// TODO Remove the conditional flag check when https://github.com/spf13/cobra/issues/655 has been resolved
		if viper.GetString("cluster") == "" {
			return fmt.Errorf(`required flag(s) "cluster" not set`)
		}
		if err := config.Load(viper.GetString("cluster"), viper.GetString("machines")); err != nil {
			return err
		}
	}
	SetLogLevel()

	if err := util.CreateDirectory(config.Cfg.TempDir, "0750"); err != nil {
		return err
	}

	if prepareSaltSSH {
		if err := salt.PrepareSaltSSH(); err != nil {
			return err
		}
	}
	return nil
}
