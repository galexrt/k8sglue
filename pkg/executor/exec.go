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

package executor

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/coreos/pkg/capnslog"
)

var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/pkg/executor", "executor")

// ExecStdoutByte executes a command and returns the output as a []byte
func ExecStdoutByte(cmdName string, args []string) ([]byte, error) {
	logger.Infof("ExecStdoutByte: Running \"%s %s\"", cmdName, strings.Join(args, " "))
	cmd := exec.Command(cmdName, args...)

	out, err := cmd.Output()
	if err != nil {
		return []byte{}, fmt.Errorf("error getting output for cmd, %+v", err)
	}

	return out, nil
}

// ExecOutToLog executes a command and logs StdOut and StdErr of the command to console (var logger)
func ExecOutToLog(logPrefix, cmdName string, args []string) error {
	logger.Infof("ExecOutToLog: Running \"%s %s\" as %s", cmdName, strings.Join(args, " "), logPrefix)
	cmd := exec.Command(cmdName, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating StdoutPipe for cmd, %+v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating StderrPipe for cmd, %+v", err)
	}

	go readerToLog(stdout, logPrefix)
	go readerToLog(stderr, logPrefix)

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting cmd, %+v", err)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("error waiting for cmd, %+v", err)
	}
	return nil
}

func readerToLog(r io.ReadCloser, prefix string) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		logger.Infof("%s | %s", prefix, sc.Text())
	}
	r.Close()
}
