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

package terraform

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/executor"
)

var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/pkg/terraform", "terraform")

// TerraformServersOutput Terraform servers JSON variable output structure
type TerraformServersOutput struct {
	Sensitive bool                        `json:"sensitive"`
	Type      string                      `json:"type"`
	Value     TerraformServersOutputValue `json:"value"`
}

// TerraformServersOutputValue variables inside the TerraformServersOutput output struct
type TerraformServersOutputValue struct {
	IDs         []string `json:"ids"`
	Names       []string `json:"names"`
	AddressIPv4 []string `json:"addresses_ipv4"`
	AddressIPv6 []string `json:"addresses_ipv6"`
}

// ApplyFull Run `terraform init`, `terraform plan` and `terraform apply` with
// auto approve and put the plan in a temp directory.
func ApplyFull() error {
	defer func() {
		logger.Info("Apply() defer called")
	}()
	logger.Info("Apply() called")

	tmp, err := ioutil.TempDir(os.TempDir(), "k8sglue")
	if err != nil {
		return err
	}

	commands := [][]string{
		{"terraform init", "terraform", "init", "-input=false"},
		{"terraform plan", "terraform", "plan", fmt.Sprintf("-out=%s/tfplan", tmp), "-input=false"},
		{"terraform apply", "terraform", "apply", "-input=false", "-auto-approve", fmt.Sprintf("%s/tfplan", tmp)},
	}

	for _, command := range commands {
		if err := executor.ExecOutToLog(command[0], command[1], command[2:]); err != nil {
			return err
		}
	}

	return nil
}

// DestroyFull run `terraform init` and `terraform destroy` with auto approve
func DestroyFull() error {
	tmp, err := ioutil.TempDir(os.TempDir(), "k8sglue")
	if err != nil {
		return err
	}
	defer func(path string) {
		os.RemoveAll(path)
	}(tmp)

	commands := [][]string{
		{"terraform init", "terraform", "init", "-input=false"},
		{"terraform destroy", "terraform", "destroy", "-auto-approve"},
	}

	for _, command := range commands {
		if err := executor.ExecOutToLog(command[0], command[1], command[2:]); err != nil {
			return err
		}
	}

	return nil
}

// GetOutput gets a terraform output by name
func GetOutput(name string) ([]byte, error) {
	commands := [][]string{
		{"terraform output", "terraform", "output", "-json", name},
	}

	var out []byte
	var err error
	for _, command := range commands {
		if out, err = executor.ExecStdoutByte(command[1], command[2:]); err != nil {
			return nil, err
		}
	}

	return out, nil
}
