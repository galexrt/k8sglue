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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/executor"
)

// TODO Copy the actual public key from the host0
// TODO Look into using a simpler if key exists check
var saltKeyAcceptMagic = `if salt-key --out=pprint -f '%s' | grep -q '%s'; then
	salt-key --out=json -q -y -a %s
else
	echo "Key not found on this master"
	exit 1
fi`

// KeyAccept use salt-ssh to get the fingerprint of the minion from **one** server
// and then accept it on each salt-master(s)
func KeyAccept(hostname string) error {
	outTempFile, err := ioutil.TempFile(os.TempDir(), "k8sglue")
	if err != nil {
		return err
	}
	if err = outTempFile.Close(); err != nil {
		return err
	}

	args := append(getSaltSSHDefaultArgs(),
		"--out=json",
		"--static",
		fmt.Sprintf("--out-file=%s", outTempFile.Name()),
		hostname,
		"cmd.run",
		"salt-call --out=json --local key.finger",
	)

	if err = executor.ExecOutToLog("salt-ssh key.finger", SaltSSHCommand, args); err != nil {
		return err
	}

	out, err := ioutil.ReadFile(outTempFile.Name())
	if err != nil {
		return err
	}

	outParsed := make(map[string]string, 1)
	if err = json.Unmarshal(out, &outParsed); err != nil {
		return err
	}

	returnParsed := make(map[string]string, 1)
	if err = json.Unmarshal([]byte(outParsed[hostname]), &returnParsed); err != nil {
		return err
	}

	fingerprint := returnParsed["local"]
	if len(fingerprint) == 0 {
		return fmt.Errorf("salt-minion on %s did not return key fingerprint", hostname)
	}

	mastersRoster := config.Cfg.Cluster.Salt.Masters.GetEntriesByRole("salt-master")
	if len(mastersRoster) == 0 {
		return fmt.Errorf("no nodes with role salt-master found")
	}

	var masters string
	for master := range mastersRoster {
		masters += master + ","
	}
	masters = strings.TrimRight(masters, ",")

	args = append(getSaltSSHDefaultArgs(),
		"-L",
		masters,
		"cmd.run",
		fmt.Sprintf(saltKeyAcceptMagic, hostname, fingerprint, hostname),
	)

	return executor.ExecOutToLog("salt-ssh salt-key", SaltSSHCommand, args)
}

// KeyAcceptList parallel loop over a list of machines and running KeyAccept on it
func KeyAcceptList() error {
	return nil
}
