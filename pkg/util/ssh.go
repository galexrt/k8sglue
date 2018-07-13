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

package util

import (
	"fmt"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/models"
	"github.com/galexrt/k8sglue/pkg/sshexecutor"
)

var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/pkg/util", "ssh")

// GetSSHExecutorForMachine return started sshexecutor.SSHExecutor struct from models.Machine
func GetSSHExecutorForMachine(machine models.Machine) (*sshexecutor.SSHExecutor, error) {
	if machine.SSHPort == 0 {
		if config.Cfg.Cluster.Cluster.SSHPort != 0 {
			machine.SSHPort = config.Cfg.Cluster.Cluster.SSHPort
		} else {
			logger.Warningf("no port for machine and cluster config given, defaulting to %d", sshexecutor.DefaultSSHPort)
			machine.SSHPort = sshexecutor.DefaultSSHPort
		}
	}

	mergedSSHConfig, err := machine.MergeSSHConfig(config.Cfg.Cluster.Cluster.SSHConfig)
	if err != nil {
		return nil, err
	}
	logger.Debugf("ssh.ClientConfig for machine %s: %+v", machine.Hostname, mergedSSHConfig)

	host := fmt.Sprintf("%s:%d", machine.Hostname, machine.SSHPort)
	sshExec := sshexecutor.New(host, mergedSSHConfig)
	logger.Debugf("starting connection infos for %s", machine.Hostname)
	err = sshExec.Start()
	return sshExec, err
}
