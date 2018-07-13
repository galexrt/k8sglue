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
	"os"
	"path"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/tar"
)

// Tar Create salt states directory
func Tar() error {
	saltTar, err := os.OpenFile(
		path.Join(config.Cfg.TempDir, "salt.tar.gz"),
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		os.FileMode(0660),
	)
	if err != nil {
		return err
	}
	defer saltTar.Close()
	return tar.CreateTarGz("salt/", saltTar)
}
