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

import (
	"encoding/json"
	"fmt"

	"github.com/coreos/pkg/capnslog"
	"github.com/galexrt/k8sglue/pkg/config"
	"github.com/galexrt/k8sglue/pkg/terraform"
	"github.com/spf13/viper"
)

type Cluster struct {
	Cluster *config.Cluster
	Nodes   []config.Node
}

var logger = capnslog.NewPackageLogger("github.com/galexrt/k8sglue/pkg/cluster", "cluster")

func ValidateCluster(clusterFile string) (*config.Cluster, error) {
	viper.SetConfigFile(clusterFile)

	viper.AutomaticEnv()

	// If a cluster file is found, read it in.
	logger.Infof("using cluster file: %s", viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("unable to read cluster, %+v", err)
	}

	clusterCfg := &config.Cluster{}
	if err := viper.Unmarshal(clusterCfg); err != nil {
		return nil, fmt.Errorf("unable to decode into cluster struct, %v", err)
	}

	return clusterCfg, nil
}

func New(clusterCfg *config.Cluster) *Cluster {
	return &Cluster{
		Cluster: clusterCfg,
		Nodes:   []config.Node{},
	}
}

func (c *Cluster) GetNodeList() ([]config.Node, error) {
	nodes := []config.Node{}

	tfNodes, err := c.GetTerraformNodeList()
	if err != nil {
		return nodes, err
	}

	return append(tfNodes, c.Cluster.Machines.ExternalNodes...), nil
}

func (c *Cluster) GetTerraformNodeList() ([]config.Node, error) {
	nodes := []config.Node{}

	tfOutput, err := terraform.GetOutput("servers")
	if err != nil {
		return nodes, err
	}

	servers := &TerraformServersOutput{}
	err = json.Unmarshal(tfOutput, servers)
	if err != nil {
		return nodes, err
	}

	for i := range servers.Value.IDs {
		nodes = append(nodes, config.Node{
			ID:       servers.Value.IDs[i],
			Hostname: servers.Value.Names[i],
			Addresses: config.Addresses{
				IPv4: []string{
					servers.Value.AddressIPv4[i],
				},
				IPv6: []string{
					servers.Value.AddressIPv6[i],
				},
			},
			Roles: config.Roles{},
		})
	}

	return nodes, nil
}
