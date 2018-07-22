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
	"time"
)

// Certs run all commands necessary to either renew or create new certificates for the salt-master(s) + sync them.
func Certs(targets, names []string, validFor time.Duration, renewTime time.Duration, dryrun bool) error {
	_, _, err := CertsCheck(targets)
	if err != nil {
		return err
	}

	cert, key, err := CertsGenerate(names, validFor)
	if err != nil {
		return err
	}

	if err = CertsSync(targets, cert, key); err != nil {
		return err
	}

	return nil
}
