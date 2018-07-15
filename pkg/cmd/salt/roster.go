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

package salt

import (
	"github.com/galexrt/k8sglue/pkg/config"
)

// Roster prints out the master roster file
func Roster() ([]byte, error) {
	opts := map[string]interface{}{
		"grains": map[string]interface{}{
			"roles": []string{
				"salt-master",
			},
		},
	}
	if err := config.Cfg.Cluster.Salt.Masters.AddMinionOpts(opts, true); err != nil {
		return nil, err
	}
	return config.Cfg.Cluster.Salt.Masters.ToByte()
}
