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
	"io/ioutil"
	"os"
	"path"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/models"
	yaml "gopkg.in/yaml.v2"
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
	/* TODO Uncomment when testing is done
	tempDir, err := ioutil.TempDir(os.TempDir(), appName)
	if err != nil {
		return nil
	}*/
	tempDir := "/tmp/k8sglue"

	startDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// TODO make configurable by flag
	saltDir := path.Join(startDir, "salt")

	Cfg = &Config{
		Cluster:  &models.Cluster{},
		LogLevel: capnslog.INFO,
		SaltDir:  saltDir,
		TempDir:  tempDir,
		StartDir: startDir,
	}
	return nil
}

// Load load cluster config into Cfg variable
func Load(configPath string) error {
	cluster, err := LoadCluster(configPath)
	cluster.Salt.DefaultRosterData.Host = ""
	Cfg.Cluster = cluster
	return err
}

// LoadCluster load a cluster config
func LoadCluster(configPath string) (*models.Cluster, error) {
	out, err := loadYAML(configPath)
	if err != nil {
		return nil, err
	}
	cluster := &models.Cluster{}
	if err := yaml.Unmarshal(out, cluster); err != nil {
		return nil, err
	}
	return cluster, nil
}

func loadYAML(configPath string) ([]byte, error) {
	return ioutil.ReadFile(configPath)
}
