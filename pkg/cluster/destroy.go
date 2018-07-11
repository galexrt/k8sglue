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

package doer

import "github.com/galexrt/k8sglue/pkg/terraform"

func (c *Cluster) Destroy() error {
	defer func() {
		logger.Info("Destroy() defer called")
	}()
	logger.Info("Destroy() called")

	if c.Cluster.Machines.Terraform.Enabled {
		if err := terraform.DestroyFull(); err != nil {
			return err
		}
	}

	return nil
}
