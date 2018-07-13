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
	"time"

	"github.com/galexrt/k8sglue/pkg/cert"
	"github.com/galexrt/k8sglue/pkg/config"
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
func CertsCopy(saltCertPath, saltKeyPath string) error {
	// TODO Add functionality to salt.CertsCopy() which copies the certificates to the machines

	return nil
}
