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

package sshexecutor

import (
	"fmt"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/models"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// GetSSHForMachine return started SSHExecutor struct from models.Machine
func GetSSHForMachine(machine models.Machine) (*SSHExecutor, error) {
	mergedSSHConfig, err := getSSHConnectionDetails(machine)
	if err != nil {
		return nil, err
	}

	host := fmt.Sprintf("%s:%d", machine.Hostname, machine.SSHPort)
	sshExec := New(host, mergedSSHConfig)
	logger.Debugf("starting connection to %s", machine.Hostname)
	return sshExec, sshExec.Start()
}

// GetSFTPForMachine return sftp.Client from a SSHExecutor ssh.Client connection
func GetSFTPForMachine(machine models.Machine) (*sftp.Client, error) {
	sshExec, err := GetSSHForMachine(machine)
	if err != nil {
		return nil, err
	}
	return sshExec.SFTP()
}

func getSSHConnectionDetails(machine models.Machine) (*ssh.ClientConfig, error) {
	if machine.SSHPort == 0 {
		if config.Cfg.Cluster.Cluster.SSHPort != 0 {
			machine.SSHPort = config.Cfg.Cluster.Cluster.SSHPort
		} else {
			logger.Warningf("no port for machine and cluster config given, defaulting to %d", DefaultSSHPort)
			machine.SSHPort = DefaultSSHPort
		}
	}

	mergedSSHConfig, err := machine.MergeSSHConfig(config.Cfg.Cluster.Cluster.SSHConfig)
	logger.Debugf("ssh.ClientConfig for machine %s: %+v", machine.Hostname, mergedSSHConfig)
	return mergedSSHConfig, err
}
