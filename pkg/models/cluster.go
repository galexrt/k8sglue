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

package models

import (
	saltmodels "github.com/galexrt/k8sglue/pkg/salt/models"
)

// Cluster holds special cluster information.
type Cluster struct {
	Salt       Salt       `yaml:"salt"`
	Hostnames  Hostnames  `yaml:"hostnames"`
	Kubernetes Kubernetes `yaml:"kubernetes"`
}

// Salt will hold all required information for cluster setup.
type Salt struct {
	DefaultRosterData saltmodels.RosterData `yaml:"defaultRosterData"`
}

// Hostnames holds info about which hostnames/addresses to use for certain services
type Hostnames struct {
	Salt       HostnamesSalt       `yaml:"salt"`
	Kubernetes HostnamesKubernetes `yaml:"kubernetes"`
}

// HostnamesSalt
type HostnamesSalt struct {
}

// HostnamesKubernetes
type HostnamesKubernetes struct {
}

// Kubernetes will hold all Kubernetes and kubeadm related settings.
type Kubernetes struct {
	Kubeadm Kubeadm `yaml:"kubeadm"`
}

// Kubeadm contains kubeadm configurations which will be used for the cluster.
type Kubeadm struct {
	// TODO Add kubeadm settings, which is used on every node (Kubernetes master and worker)
}
