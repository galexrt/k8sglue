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
	"github.com/imdario/mergo"
	"golang.org/x/crypto/ssh"
)

// MachineList a list of Machines.
type MachineList struct {
	// Machines list
	Machines []Machine `yaml:"machines"`
}

// Machine holds information about a machine/host/node.
type Machine struct {
	Hostname  string            `yaml:"hostname"`
	SSHConfig *ssh.ClientConfig `yaml:"sshConfig"`
	SSHPort   uint16            `yaml:"sshPort"`
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
	ETCD *bool `yaml:"etcd"`
	// Kubernetes roles
	Kubernetes *RolesKubernetes `yaml:"kubernetes"`
	// Salt roles
	Salt *RolesSalt `yaml:"salt"`
}

// RolesKubernetes contains the Kubernetes roles a machine/host/node can have.
type RolesKubernetes struct {
	// Kubernetes Master role
	Master *bool `yaml:"master"`
	// Kubernetes Worker role
	Worker *bool `yaml:"worker"`
}

// RolesSalt contains the Salt roles a machine/host/node can have.
type RolesSalt struct {
	// Salt Master role
	Master *bool `yaml:"master"`
}

// MergeSSHConfig merges a given ssh.ClientConfig with the machine's SSHConfig ssh.ClientConfig
func (m *Machine) MergeSSHConfig(global *ssh.ClientConfig) (*ssh.ClientConfig, error) {
	if m.SSHConfig == nil {
		m.SSHConfig = global
	}
	if err := mergo.Merge(m.SSHConfig, *global); err != nil {
		return nil, err
	}
	return m.SSHConfig, nil
}

// GetMachineByHostname get a machine by hostname
func (ml *MachineList) GetMachineByHostname(hostname string) *Machine {
	for _, machine := range ml.Machines {
		if machine.Hostname == hostname {
			return &machine
		}
	}
	return nil
}

// GetMachinesByRole get a machine by matching roles
func (ml *MachineList) GetMachinesByRole(roles Roles) []Machine {
	var machines []Machine
	for _, machine := range ml.Machines {
		if roles.ETCD != nil {
			if machine.Roles.ETCD != nil &&
				*machine.Roles.ETCD == *roles.ETCD {
				machines = append(machines, machine)
			}
		}
		if roles.Kubernetes != nil &&
			machine.Roles.Kubernetes != nil {
			if roles.Kubernetes.Master != nil &&
				machine.Roles.Kubernetes.Master != nil &&
				*machine.Roles.Kubernetes.Master == *roles.Kubernetes.Master {
				machines = append(machines, machine)
			}
			if roles.Kubernetes.Worker != nil &&
				machine.Roles.Kubernetes.Worker != nil &&
				*machine.Roles.Kubernetes.Worker == *roles.Kubernetes.Worker {
				machines = append(machines, machine)
			}
		}
		if roles.Salt != nil && machine.Roles.Salt != nil {
			if roles.Salt.Master != nil &&
				machine.Roles.Salt.Master != nil &&
				*machine.Roles.Salt.Master == *roles.Salt.Master {
				machines = append(machines, machine)
			}
		}
	}
	return machines
}
