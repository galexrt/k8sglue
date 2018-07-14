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
	"os"
	"path"
	"sync"
	"time"

	"github.com/galexrt/k8sglue/pkg/cert"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/models"
	"github.com/galexrt/k8sglue/pkg/sshexecutor"
)

// Certs generate salt-master certs
func Certs(saltMasterIPs []string) (string, string, error) {
	saltCert, saltKey, err := cert.Generate(saltMasterIPs, "", 24*time.Hour, false, 4096, "P521")
	if err != nil {
		return "", "", err
	}

	certPath := path.Join(config.Cfg.TempDir, "salt-master-cert.pem")
	cert, err := os.OpenFile(
		certPath,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		os.FileMode(0660),
	)
	if err != nil {
		return "", "", err
	}
	defer cert.Close()

	var written int
	if written, err = cert.Write(saltCert); err != nil {
		return "", "", err
	}
	if written != len(saltCert) {
		return "", "", fmt.Errorf("salt-master cert didn't write all bytes (cert: %d, written: %d)", len(saltCert), written)
	}

	keyPath := path.Join(config.Cfg.TempDir, "salt-master-key.pem")
	key, err := os.OpenFile(
		keyPath,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		os.FileMode(0660),
	)
	if err != nil {
		return "", "", err
	}
	defer key.Close()

	if written, err = key.Write(saltKey); err != nil {
		return "", "", err
	}
	if written != len(saltKey) {
		return "", "", fmt.Errorf("salt-master key didn't write all bytes (key: %d, written: %d)", len(saltKey), written)
	}

	return certPath, keyPath, nil
}

// CertsCopy copy the certs by path to the salt-master(s) machines
func CertsCopy(masters []models.Machine, saltCertPath, saltKeyPath string) error {
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
	return nil
}
