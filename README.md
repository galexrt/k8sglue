# k8sglue

This project is a total work in progress right now!

## Requirements

* Saltstack (`salt-ssh`)
* Golang for compiling the project (tested with >= `1.10`)

## Goals

### Primary Goals

* Setup salt-master(s) infrastructure.
    * Deployment and management of nodes.
    * Used for salt beacons to monitor nodes (though the actual monitoring is secondary).
* Kubernetes cluster(s) managed with Saltstack.
    * With `kubeadm` under the hood.

### Secondary Goals

* Salt Beacons should be implemented for additional monitoring and reacting: "automatic problem solver" (e.g. a machine runs out of memory = run `logrotate`, a machine is not ready with kubelet PLEG errors = reboot machine).
* Servers are host firewalled (with exceptions, e.g. allow custom which allows certain ports).
    * Either by Saltstack or depending on how "good" Canal (Calico) allows to "host firewall" using `GlobalNetworkPolicy`.

## What software is used for what

* k8sglue
    * Generate Salt SSH Roster file.
    * Triggers salt-ssh to install salt-master(s).
* Saltstack
    * "Actual" configuration of servers.

## Salt

For information on the Salt part, check out the [salt/README.md](/salt/README.md).
