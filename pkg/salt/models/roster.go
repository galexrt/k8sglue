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
	"fmt"

	"github.com/galexrt/k8sglue/pkg/util"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

// Roster the format of a Saltstack Roster file as seen here: https://docs.saltstack.com/en/2017.7/topics/ssh/roster.html#targets-data
type Roster map[string]*RosterData

// RosterData a roster host entry as seen here: https://docs.saltstack.com/en/2017.7/topics/ssh/roster.html#targets-data
type RosterData struct {
	Host       string                 `yaml:"host,omitempty"`
	User       string                 `yaml:"user,omitempty"`
	Passwd     string                 `yaml:"passwd,omitempty"`
	Port       int16                  `yaml:"port,omitempty"`
	Sudo       bool                   `yaml:"sudo,omitempty"`
	SudoUser   string                 `yaml:"sudo_user,omitempty"`
	TTY        bool                   `yaml:"tty,omitempty"`
	Priv       string                 `yaml:"priv,omitempty"`
	Timeout    string                 `yaml:"timeout,omitempty"`
	MinionOpts map[string]interface{} `yaml:"minion_opts,omitempty"`
	ThinDir    string                 `yaml:"thin_dir,omitempty"`
	CMDUmask   uint8                  `yaml:"cmd_umask,omitempty"`
}

// Merge merge Roster together
func (r Roster) Merge(src Roster) error {
	for k, v := range src {
		if _, ok := r[k]; ok {
			return fmt.Errorf("machine %s already in roster", k)
		}
		r[k] = v
	}
	return nil
}

// GetHosts returns all `RosterData.Host` names in a []string
func (r Roster) GetHosts() []string {
	hosts := []string{}
	for host := range r {
		hosts = append(hosts, host)
	}
	return hosts
}

// GetEntriesByRole get all entries that have a certain role
func (r Roster) GetEntriesByRole(role string) Roster {
	entries := Roster{}
	for k, entry := range r {
		if checkForRole(entry.MinionOpts, role) {
			entries[k] = entry
		}
	}
	return entries
}

func checkForRole(minionOpts map[string]interface{}, role string) bool {
	if _, ok := minionOpts["grains"]; ok {
		grains := util.ConvertMapInterfaceToMapStringInterface(minionOpts["grains"].(map[interface{}]interface{}))
		if _, ok = grains["roles"]; ok {
			roles := util.ConvertInterfaceSliceToStringSlice(grains["roles"].([]interface{}))
			for _, r := range roles {
				if r == role {
					return true
				}
			}
		}
	}
	return false
}

// AddMinionOpts add minion_opts to all RosterData entries
func (r Roster) AddMinionOpts(opts map[string]interface{}, overwrite bool) error {
	for k := range r {
		if len(r[k].MinionOpts) == 0 {
			r[k].MinionOpts = opts
			continue
		}
		var err error
		if overwrite {
			err = mergo.Map(r[k].MinionOpts, opts, mergo.WithOverride)
		} else {
			err = mergo.Map(r[k].MinionOpts, opts)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ToByte marshals the structure into YAML for salt to use it
func (r Roster) ToByte() ([]byte, error) {
	return yaml.Marshal(r)
}

// ToByte marshals the structure into YAML for salt to use it
func (rd RosterData) ToByte() ([]byte, error) {
	return yaml.Marshal(rd)
}

// Merge merge Roster together
func (rd RosterData) Merge(src RosterData) error {
	return mergo.MergeWithOverwrite(&rd, src)
}

// SetDefaultRosterData sets these defaults on all machines if not set
func (r Roster) SetDefaultRosterData(src RosterData) error {
	for k := range r {
		if err := mergo.MergeWithOverwrite(r[k], src); err != nil {
			return err
		}
	}
	return nil
}
