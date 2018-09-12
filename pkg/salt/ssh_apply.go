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
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/executor"
)

// HighState definition of high state for `state.apply` call here
const HighState = ""

// SSHApply trigger salt-ssh highstate using salt-ssh on the salt-master(s)
func SSHApply(machines []string, slsFiles string) error {
	sort.Strings(machines)
	args := append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh", "state.apply")
	if err := executor.ExecOutToLog("salt-ssh state.apply all master", SaltSSHCommand, args); err != nil {
		return err
	}

	for _, machine := range machines {
		args = append(getSaltSSHDefaultArgs(),
			generateTargetFlags([]string{machine})...,
		)
		args = append(args, "--state-verbose=false", "--refresh", "--static", "--out=json", "cp.get_file_str", "/etc/salt/pki/minion/minion.pub")
		out, err := executor.ExecStdoutByte(SaltSSHCommand, args)
		if err != nil {
			return err
		}
		var parsed map[string]string
		if err = json.Unmarshal(out, &parsed); err != nil {
			return err
		}
		if err = ioutil.WriteFile(
			path.Join(config.Cfg.TempDir, "data", fmt.Sprintf("pub-%s", machine)),
			[]byte(parsed[machine]),
			0600,
		); err != nil {
			return err
		}
	}

	for _, machine := range machines {
		args = append(getSaltSSHDefaultArgs(),
			generateTargetFlags(machines)...,
		)
		args = append(args, "--state-verbose=false", "--refresh",
			"state.single",
			"file.managed",
			fmt.Sprintf("name=%s", path.Join("/etc/salt/pki/master/minions", machine)),
			fmt.Sprintf("source=salt://pub-%s", machine),
			"dir_mode=0700",
			"mode=0600",
			"user=root",
			"group=root",
		)
		if err := executor.ExecOutToLog("salt-ssh copy master minion pub keys", SaltSSHCommand, args); err != nil {
			return err
		}

		args = append(getSaltSSHDefaultArgs(),
			generateTargetFlags(machines)...,
		)
		args = append(args, "--state-verbose=false", "--refresh",
			"state.single",
			"cmd.run",
			fmt.Sprintf("name=rm -f /etc/salt/pki/master/minions_pre/%s", machine),
		)
		if err := executor.ExecOutToLog("salt-ssh remove master pre acc keys", SaltSSHCommand, args); err != nil {
			return err
		}
	}

	files := []string{
		"/etc/salt/pki/master/master.pem",
		"/etc/salt/pki/master/master.pub",
		"/etc/salt/pki/master/master_sign.pem",
		"/etc/salt/pki/master/master_sign.pub",
		"/etc/salt/pki/master/master_pubkey_signature",
	}
	for _, file := range files {
		args = append(getSaltSSHDefaultArgs(),
			generateTargetFlags(machines[0:0])...,
		)
		args = append(args, "--state-verbose=false", "--refresh", "--static", "--out=json", "cp.get_file_str", file)
		out, err := executor.ExecStdoutByte(SaltSSHCommand, args)
		if err != nil {
			return err
		}
		var parsed map[string]string
		if err = json.Unmarshal(out, &parsed); err != nil {
			return err
		}

		if err = ioutil.WriteFile(
			path.Join(config.Cfg.TempDir, "data", filepath.Base(file)),
			[]byte(parsed[machines[0]]),
			0600,
		); err != nil {
			return err
		}
	}

	for _, file := range files {
		args = append(getSaltSSHDefaultArgs(),
			generateTargetFlags(machines)...,
		)
		args = append(args, "--state-verbose=false", "--refresh",
			"state.single",
			"file.managed",
			fmt.Sprintf("name=%s", file),
			fmt.Sprintf("source=salt://%s", filepath.Base(file)),
			"dir_mode=0700",
			"mode=0600",
			"user=root",
			"group=root",
		)
		if err := executor.ExecOutToLog("salt-ssh copy master sign files", SaltSSHCommand, args); err != nil {
			return err
		}
	}
	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh",
		"state.single",
		"file.managed",
		fmt.Sprintf("name=%s", strings.Replace("/etc/salt/pki/master/master_sign.pub", "master", "minion", 1)),
		fmt.Sprintf("source=salt://%s", filepath.Base("/etc/salt/pki/master/master_sign.pub")),
		"dir_mode=0700",
		"mode=0600",
		"user=root",
		"group=root",
	)
	if err := executor.ExecOutToLog("salt-ssh copy master sign files", SaltSSHCommand, args); err != nil {
		return err
	}

	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh",
		"state.single",
		"file.managed",
		"name=/etc/salt/ssh/id_rsa",
		"source=salt://ssh_id_rsa",
		"dir_mode=0700",
		"mode=0600",
		"user=root",
		"group=root",
	)

	if err := executor.ExecOutToLog("salt-ssh copy ssh key", SaltSSHCommand, args); err != nil {
		return err
	}

	if err := Sync(machines); err != nil {
		return err
	}

	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh",
		"state.single",
		"file.managed",
		"name=/srv/pillar/salt_master_addresses.yaml",
		"source=salt://salt_master_addresses.yaml",
		"dir_mode=0700",
		"mode=0600",
		"user=root",
		"group=root",
	)

	if err := executor.ExecOutToLog("salt-ssh copy salt_master_addresses.yaml", SaltSSHCommand, args); err != nil {
		return err
	}

	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh", "state.apply")
	if err := executor.ExecOutToLog("salt-ssh state.apply all master", SaltSSHCommand, args); err != nil {
		return err
	}

	args = append(getSaltSSHDefaultArgs(),
		generateTargetFlags(machines)...,
	)
	args = append(args, "--state-verbose=false", "--refresh",
		"state.single",
		"cmd.run",
		"name=systemctl restart salt-master salt-minion",
	)

	return executor.ExecOutToLog("salt-ssh restart salt-minion", SaltSSHCommand, args)
}
