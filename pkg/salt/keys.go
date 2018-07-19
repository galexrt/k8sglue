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

// TODO Add retry logic for the bash script
var saltKeyAcceptMagic = `machines=(
%s
)
fingerprints=(
%s
)
ERROR=0
for (( i = 0; i < ${#machines[@]}; ++i)); do
	if salt-key --out=pprint -f "${machines[i]}" | grep -q "${fingerprints[i]}"; then
		echo "Accepting key of ${machines[i]}"
		salt-key --out=json -q -y -a "${machines[i]}"
	else
		echo "Key for minion ${machines[i]} not found on this master."
		ERROR=1
	fi
done
if [ "${ERROR}" = "1" ]; then
	echo "At least one minion failed to be found on this master."
	exit 1
fi`

// KeyAccept use salt-ssh to get the fingerprint of the minion from **one** server
// and then accept it on each salt-master(s)
func KeyAccept(machines []string) error {
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
		"--list",
		strings.Join(machines, ","),
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

	fingerprints := map[string]string{}

	for host, rawReturn := range outParsed {
		returnParsed := make(map[string]string, 1)
		if err = json.Unmarshal([]byte(rawReturn), &returnParsed); err != nil {
			return err
		}
		fingerprint := returnParsed["local"]
		if len(fingerprint) == 0 {
			return fmt.Errorf("salt-minion on %s did not return key fingerprint", host)
		}
		fingerprints[host] = fingerprint
	}

	hostsArray := []string{}
	fingerprintsArray := []string{}
	for host, fingerprint := range fingerprints {
		hostsArray = append(hostsArray, host)
		fingerprintsArray = append(fingerprintsArray, fingerprint)
	}

	mastersRoster := config.Cfg.Machines.GetEntriesByRole("salt-master")
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
		fmt.Sprintf(saltKeyAcceptMagic, strings.Join(hostsArray, "\n"), strings.Join(fingerprintsArray, "\n")),
	)

	return executor.ExecOutToLog("salt-ssh salt-key", SaltSSHCommand, args)
}

// KeyRemove removes minions keys from salt-master(s)
func KeyRemove(machines []string) error {
	// TODO Implement functionality
	return nil
}
