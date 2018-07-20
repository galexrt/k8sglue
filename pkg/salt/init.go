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
	"fmt"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/spf13/viper"
)

var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/pkg/salt", "salt")

// Init call steps necessary to init salt-master(s)
func Init(machines []string) error {
	hosts := viper.GetStringSlice("host")
	if len(hosts) == 0 && !viper.GetBool("all") {
		return fmt.Errorf("no all or host flag given")
	} else if viper.GetBool("all") {
		hosts = config.Cfg.Machines.GetHosts()
	}

	if err := Ping(hosts); err != nil {
		return err
	}

	if _, _, err := Secrets(); err != nil {
		return err
	}

	if err := Apply([]string{}, HighState); err != nil {
		return err
	}

	if err := Sync(); err != nil {
		return err
	}

	return nil
}
