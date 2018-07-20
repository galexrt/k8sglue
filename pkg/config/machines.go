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
	"path/filepath"

	saltmodels "github.com/galexrt/k8sglue/pkg/salt/models"
	yaml "gopkg.in/yaml.v2"
)

// Machines holds all the machines from the machines config files and has some nice functions for that
type Machines struct {
	Roster saltmodels.Roster
}

// LoadMachines load all machines files
func LoadMachines(globPath string) (*saltmodels.Roster, error) {
	machines := &saltmodels.Roster{}
	paths, err := filepath.Glob(globPath)
	if err != nil {
		return nil, nil
	}
	for _, file := range paths {
		out, err := loadYAML(file)
		if err != nil {
			return nil, err
		}
		roster := saltmodels.Roster{}
		if err = yaml.Unmarshal(out, &roster); err != nil {
			return nil, err
		}
		if err = machines.Merge(roster); err != nil {
			return nil, err
		}
	}

	return machines, nil
}