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

// MachineList a list of Machines.
type MachineList struct {
	// Machines list
	Machines []Machine `yaml:"machines"`
}

// Machine holds information about a machine/host/node.
type Machine struct {
	Hostname string `yaml:"hostname"`
	// (Unique) Server ID
	ID        string    `yaml:"id"`
	Addresses Addresses `yaml:"addresses"`
	Roles     Roles     `yaml:"roles"`
	// Parameters will be added to the Salt Roster host entry
	Parameters map[string]interface{} `yaml:"parameters"`
}

// Addresses holds IPv4 and IPv6 address lists of a machine/host/node.
type Addresses struct {
	IPv4 []string `yaml:"ipv4"`
	IPv6 []string `yaml:"ipv6"`
}

// Roles contains the roles a machine/host/node can have.
type Roles struct {
	// ETCD role
	ETCD bool `yaml:"etcd"`
	// Kubernetes roles
	Kubernetes RolesKubernetes `yaml:"kubernetes"`
	// Salt roles
	Salt RolesSalt `yaml:"salt"`
}

// RolesKubernetes contains the Kubernetes roles a machine/host/node can have.
type RolesKubernetes struct {
	// Kubernetes Master role
	Master bool `yaml:"master"`
	// Kubernetes Worker role
	Worker bool `yaml:"worker"`
}

// RolesSalt contains the Salt roles a machine/host/node can have.
type RolesSalt struct {
	// Salt Master role
	Master bool `yaml:"master"`
}
