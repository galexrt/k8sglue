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
	"path"

	"github.com/galexrt/k8sglue/pkg/models"
	yaml "gopkg.in/yaml.v2"
)

const (
	// ClusterConfigName cluster config name
	ClusterConfigName = "cluster.yaml"
	// KubeadmConfigName kubeadm config name
	KubeadmConfigName = "kubeadm.yaml"
	// MachinesConfigName machines list config name
	MachinesConfigName = "machines.yaml"
)

// Config holds all the configs about a cluster
type Config struct {
	TempDir     string
	Cluster     *models.Cluster
	Kubeadm     *models.Kubeadm
	MachineList *models.MachineList
}

// Cfg is a Config struct pointer to be able to access all configs from "anywhere"
var Cfg *Config

// Load load configs into Cfg variable
func Load(configPath, tempDir string) error {
	cluster, err := LoadCluster(configPath)
	if err != nil {
		return err
	}
	kubeadm, err := LoadKubeadm(configPath)
	if err != nil {
		return err
	}
	machineList, err := LoadMachineList(configPath)
	if err != nil {
		return err
	}
	Cfg = &Config{
		Cluster:     cluster,
		Kubeadm:     kubeadm,
		MachineList: machineList,
		TempDir:     tempDir,
	}
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

// LoadMachineList load a machine list config
func LoadMachineList(configPath string) (*models.MachineList, error) {
	out, err := loadYAML(path.Join(configPath, MachinesConfigName))
	if err != nil {
		return nil, err
	}
	machineList := &models.MachineList{}
	if err := yaml.Unmarshal(out, machineList); err != nil {
		return nil, err
	}
	return machineList, nil
}

func loadYAML(configPath string) ([]byte, error) {
	return ioutil.ReadFile(configPath)
}
