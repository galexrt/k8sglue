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
	"sync"

	"github.com/galexrt/k8sglue/pkg/models"
	"github.com/galexrt/k8sglue/pkg/sshexecutor"
)

// Sync syncs the salt/ directory to the salt-master(s) machines
func Sync(masters []models.Machine) error {
	errs := make(chan error, 1)

	wg := sync.WaitGroup{}
	for _, machine := range masters {
		logger.Debugf("getting connection infos for %s", machine.Hostname)
		wg.Add(1)
		go func(machine models.Machine) {
			defer wg.Done()
			sftpClient, err := sshexecutor.GetSFTPForMachine(machine)
			if err != nil {
				errs <- err
				return
			}
			defer sftpClient.Close()

			if err != nil {
				errs <- err
				return
			}
			files, err := sftpClient.ReadDir(".")
			if err != nil {
				errs <- err
				return
			}
			fmt.Printf("TEST: %+v\n", files)
		}(machine)
	}
	logger.Infof("waiting for all salt-master syncs to be completed")
	wg.Wait()
	close(errs)

	var err error
	for err = range errs {
		logger.Errorf("salt sync error. %+v", err)
	}

	return err
}
