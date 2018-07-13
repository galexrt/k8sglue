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
	"github.com/coreos/pkg/capnslog"
	"golang.org/x/crypto/ssh"
)

const DefaultSSHPort = 22

var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/pkg/ssh", "ssh")

// SSHExecutor holds the info for one host connection
type SSHExecutor struct {
	Host      string
	SSHConfig *ssh.ClientConfig
	client    *ssh.Client
	session   *ssh.Session
}

// New creates a SSHExecutor struct, ready for use
func New(host string, sshCfg *ssh.ClientConfig) *SSHExecutor {
	return &SSHExecutor{
		Host:      host,
		SSHConfig: sshCfg,
	}
}

// Start starts connection to host with the given ssh.ClientConfig
func (se *SSHExecutor) Start() error {
	var err error
	se.client, err = ssh.Dial("tcp", se.Host, se.SSHConfig)
	if err != nil {
		return err
	}

	se.session, err = se.client.NewSession()
	if err != nil {
		return err
	}
	return nil
}

// ExecStdoutByte run a command on the remote side and return it's stdout as []byte
func (se *SSHExecutor) ExecStdoutByte(cmd string) ([]byte, error) {
	return se.session.Output(cmd)
}

// Close calls CLose() on the ssh.Client
func (se *SSHExecutor) Close() error {
	return se.client.Close()
}
