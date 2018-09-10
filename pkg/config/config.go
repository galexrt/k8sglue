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

package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/models"
	"github.com/galexrt/k8sglue/pkg/util"
	"github.com/spf13/viper"
)

// Config holds all the configs and a cluster if loaded
type Config struct {
	Cluster  *models.Cluster
	LogLevel capnslog.LogLevel
	StartDir string
	SaltDir  string
	TempDir  string
}

// Cfg is a Config struct pointer to be able to access all configs from "anywhere"
var Cfg *Config

// Init creates a new empty Config and "saves" it to Cfg
func Init(appName string) error {
	startDir, err := os.Getwd()
	if err != nil {
		return err
	}

	tempDir, err := util.ReturnFullPath(viper.GetString("temp-dir"))
	if err != nil {
		return err
	}
	saltDir, err := util.ReturnFullPath(viper.GetString("salt-dir"))
	if err != nil {
		return err
	}

	Cfg = &Config{
		Cluster:  &models.Cluster{},
		LogLevel: capnslog.INFO,
		StartDir: startDir,
		SaltDir:  saltDir,
		TempDir:  tempDir,
	}
	return nil
}

// Load load cluster config into Cfg variable
func Load(configPath string) error {
	cluster, err := LoadCluster(configPath)
	if err != nil {
		return err
	}
	cluster.Salt.DefaultRosterData.Host = ""
	Cfg.Cluster = cluster

	if Cfg.Cluster.SSHKey == "" {
		return fmt.Errorf("no sshKey given in cluster yaml")
	}

	return nil
}

func loadYAML(configPath string) ([]byte, error) {
	return ioutil.ReadFile(configPath)
}
