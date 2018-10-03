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

// SaltInfo holds special cluster information.
type SaltInfo struct {
	Salt   Salt   `yaml:"salt,omitempty"`
	SSHKey string `yaml:"sshKey"`
}

// Salt holds all required information for cluster setup.
type Salt struct {
	DefaultRosterDataAsBase bool                  `yaml:"defaultRosterDataAsBase"`
	DefaultRosterData       saltmodels.RosterData `yaml:"defaultRosterData"`
	Roster                  *saltmodels.Roster    `yaml:"roster"`
}
