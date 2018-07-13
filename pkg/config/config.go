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
	"path/filepath"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/models"
	"golang.org/x/crypto/ssh"
	yaml "gopkg.in/yaml.v2"
)

const (
	// ClusterConfigName cluster config name
	ClusterConfigName = "cluster.yaml"
	// KubeadmConfigName kubeadm config name
	KubeadmConfigName = "kubeadm.yaml"
	// MachinesConfigName machines list config directory
	MachinesConfigName = "machines"
)

// Config holds all the configs and a cluster if loaded
type Config struct {
	Cluster  *Cluster
	LogLevel capnslog.LogLevel
	TempDir  string
}

// Cluster configs related to a cluster
type Cluster struct {
	Cluster       *models.Cluster
	Kubeadm       *models.Kubeadm
	MachineList   *models.MachineList
	SaltStatesDir string
}

// Cfg is a Config struct pointer to be able to access all configs from "anywhere"
var Cfg *Config

// Init creates a new empty Config and "saves" it to Cfg
func Init(appName string) error {
	tempDir, err := ioutil.TempDir(os.TempDir(), appName)
	if err != nil {
		return nil
	}

	Cfg = &Config{
		Cluster:  &Cluster{},
		LogLevel: capnslog.INFO,
		TempDir:  tempDir,
	}
	return nil
}

// Load load cluster configs into Cfg variable
func Load(configPath, tempDir string) error {
	cluster, err := LoadCluster(configPath)
	if err != nil {
		return err
	}
	kubeadm, err := LoadKubeadm(configPath)
	if err != nil {
		return err
	}
	machineList, err := LoadMachineLists(configPath)
	if err != nil {
		return err
	}
	Cfg.Cluster.Cluster = cluster
	if Cfg.Cluster.Cluster.SSHConfig == nil {
		sshConfig := &ssh.ClientConfig{}
		Cfg.Cluster.Cluster.SSHConfig = sshConfig
	}
	Cfg.Cluster.Cluster.SSHConfig.SetDefaults()
	Cfg.Cluster.Kubeadm = kubeadm
	Cfg.Cluster.MachineList = machineList
	return nil
}

// LoadCluster load a cluster config
func LoadCluster(configPath string) (*models.Cluster, error) {
	out, err := loadYAML(path.Join(configPath, ClusterConfigName))
	if err != nil {
		return nil, err
	}
	cluster := &models.Cluster{}
	if err := yaml.Unmarshal(out, cluster); err != nil {
		return nil, err
	}
	return cluster, nil
}

// LoadKubeadm load a kubeadm config
func LoadKubeadm(configPath string) (*models.Kubeadm, error) {
	out, err := loadYAML(path.Join(configPath, KubeadmConfigName))
	if err != nil {
		return nil, err
	}
	kubeadm := &models.Kubeadm{}
	if err := yaml.Unmarshal(out, kubeadm); err != nil {
		return nil, err
	}
	return kubeadm, nil
}

// LoadMachineLists load all machine lists in the MachinesConfigName directory
func LoadMachineLists(configPath string) (*models.MachineList, error) {
	machines := []models.Machine{}
	mlFiles, err := filepath.Glob(path.Join(configPath, MachinesConfigName, "*.yaml"))
	if err != nil {
		return nil, err
	}
	for _, mlPath := range mlFiles {
		var ml *models.MachineList
		ml, err = loadMachineListConfig(mlPath)
		if err != nil {
			return nil, err
		}
		machines = append(machines, ml.Machines...)
	}
	return &models.MachineList{
		Machines: machines,
	}, nil
}

// LoadMachineList load a single machine list config
func LoadMachineList(configPath string) (*models.MachineList, error) {
	return loadMachineListConfig(path.Join(configPath, MachinesConfigName))
}

func loadMachineListConfig(filePath string) (*models.MachineList, error) {
	out, err := loadYAML(filePath)
	if err != nil {
		return nil, err
	}
	machineList := &models.MachineList{}
	err = yaml.Unmarshal(out, machineList)
	return machineList, err
}

func loadYAML(configPath string) ([]byte, error) {
	return ioutil.ReadFile(configPath)
}
