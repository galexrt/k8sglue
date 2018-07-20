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
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/util"
)

const (
	// RosterFileName salt-master(s) roster file name
	RosterFileName = "roster-salt-master"
	// SaltfileName salt-ssh config file name
	SaltfileName = "Saltfile"
)

const (
	// LogLevelCritical Salt log level critical
	LogLevelCritical = "critical"
	// LogLevelError Salt log level error
	LogLevelError = "error"
	// LogLevelWarning Salt log level warning
	LogLevelWarning = "warning"
	// LogLevelInfo Salt log level info
	LogLevelInfo = "info"
	// LogLevelDebug Salt log level debug
	LogLevelDebug = "debug"
	// LogLevelTrace Salt log level
	LogLevelTrace = "trace"
	// DefaultLogLevel Default log level for executed salt commands
	DefaultLogLevel = LogLevelWarning
)

var saltfile = `salt-ssh:
  config_dir: {{ .Config.TempDir }}/etc
`

var saltMasterConfig = `file_roots:
  base:
    - {{ .Config.SaltDir }}
    - {{ .Config.StartDir }}

root_dir: {{ .Config.TempDir }}
pidfile: {{ .Config.TempDir }}/run/salt.pid
sock_dir: {{ .Config.TempDir }}/run
cachedir: {{ .Config.TempDir }}/cache
ssh_log_file: {{ .Config.TempDir }}/logs/ssh.log
log_file: {{ .Config.TempDir }}/logs/salt.log
state_verbose: False

roster_defaults:
{{ .Additional.RosterDefaults }}
`

// GoToSaltDir chdirs into the given salt directory
func GoToSaltDir() error {
	return os.Chdir(config.Cfg.SaltDir)
}

func templateConfigFile(name, in string, additional map[string]interface{}) (string, error) {
	tmpl, err := template.New(name).Parse(in)
	if err != nil {
		return "", err
	}
	wr := new(bytes.Buffer)
	err = tmpl.Execute(wr, map[string]interface{}{
		"Config":     config.Cfg,
		"Additional": additional,
	})
	return wr.String(), err
}

// CapnslogLogLevelToSalt convert capnslog log level to salt equivalent
func CapnslogLogLevelToSalt(logLevel capnslog.LogLevel) string {
	switch logLevel {
	case capnslog.CRITICAL:
		return LogLevelCritical
	case capnslog.ERROR:
		return LogLevelError
	case capnslog.WARNING:
		return LogLevelWarning
	case capnslog.INFO:
		return LogLevelInfo
	case capnslog.DEBUG:
		return LogLevelDebug
	case capnslog.TRACE:
		return LogLevelTrace
	}
	logger.Warningf("could not convert capnslog log level to salt, defaulting to %s", DefaultLogLevel)
	return DefaultLogLevel
}

func getSaltSSHDefaultArgs() []string {
	return []string{
		//"-w",
		fmt.Sprintf("--saltfile=%s", path.Join(config.Cfg.TempDir, SaltfileName)),
		"--roster=flat",
		fmt.Sprintf("--roster-file=%s", path.Join(config.Cfg.TempDir, RosterFileName)),
		fmt.Sprintf("--log-level=%s", CapnslogLogLevelToSalt(config.Cfg.LogLevel)),
		"--ignore-host-keys",
	}
}

// PrepareSaltSSH preparse a temp directory with all info, data and config required for `salt-ssh`
func PrepareSaltSSH() error {
	if err := GoToSaltDir(); err != nil {
		return err
	}

	out, err := Roster()
	if err != nil {
		return err
	}
	rosterFilePath := path.Join(config.Cfg.TempDir, RosterFileName)
	if err = ioutil.WriteFile(
		rosterFilePath,
		out,
		0640,
	); err != nil {
		return err
	}

	rendered, err := templateConfigFile(SaltfileName, saltfile, nil)
	if err != nil {
		return err
	}
	saltfilePath := path.Join(config.Cfg.TempDir, SaltfileName)
	if err = ioutil.WriteFile(
		saltfilePath,
		[]byte(rendered),
		0640,
	); err != nil {
		return err
	}

	for _, dir := range []string{
		"cache",
		"etc",
		"logs",
		"run",
	} {
		if err = util.CreateDirectory(path.Join(config.Cfg.TempDir, dir), "0750"); err != nil {
			return err
		}
	}

	saltMasterRosterDefaults, err := config.Cfg.Cluster.Salt.DefaultRosterData.ToByte()
	if err != nil {
		return err
	}

	rendered, err = templateConfigFile("saltmasterconfig", saltMasterConfig, map[string]interface{}{
		"RosterDefaults": util.TemplateIndent(string(saltMasterRosterDefaults), 2),
	})
	if err != nil {
		return err
	}
	saltMasterConfigPath := path.Join(config.Cfg.TempDir, "etc", "master")
	return ioutil.WriteFile(
		saltMasterConfigPath,
		[]byte(rendered),
		0640,
	)
}

func generateTargetFlags(machines []string) []string {
	if len(machines) == 0 {
		return []string{
			"*",
		}
	}
	return append([]string{"-L", strings.Join(machines, ",")})
}
