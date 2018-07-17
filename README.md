# k8sglue
This project is a total work in progress right now!

## Requirements

* Saltstack (`salt-ssh`)
* Golang to compile the project (only tested with >= `1.10`)

## Flows
### `k8sglue salt init`

This command has to be run once initially so the nodes get their "deployment" information and get configured for the Kubernetes cluster.

Starting point will be that a "MachineList" is already created containing at least 1 machine with `roles.salt.master: true`.

1. k8sglue has static lists of salt-master machines (nothing else).
    1. Generate Saltstack Roster file with `minion_opts` for the salt-master(s) and put it in the tempdir.
    1. Use `salt-ssh '*' test.ping` to check the connectivity to each salt-master(s). If an error occurs, fail.
1. k8sglue pre-generates some files:
1. Generate certificate (if possible) and then distributes it among the other salt-master(s).
    1. If it already exists, just "copy" to other masters or renew if it is below X hours.
1. Generate a "deploy only" SSH key (which is used for the connection from salt-master(s) to new machines).
    1. The private key is copied to each salt-master(s).
    1. The public key is printed out for the user/admin to put on new machines.
1. k8sglue uses `salt-ssh` to:
    1. "Sync" local `salt` directory to each salt-master(s) (to `/srv/salt`).
    1. Wait for all salt-master(s) to be ready.
    1. Trigger high state.
1. k8sglue uses a `salt key accept` mechanism to get all minion fingerprints and accept them on each salt-master(s).
    1. Connect to each "new" machine to get the fingerprint.
    1. Connect to each salt-master(s) and accept that key.
* Results in: salt-master(s) ready for use and "connected to themselves".

> **NOTE** Each new node must connect to the salt-master(s).
> Each node will be "tested" with `salt.runners.manage.safe_accept` and accepted if the fingerprints match.

### `k8sglue machine join`

1. k8sglue writes a Saltstack Roster file with given roles for the node(s) given specified as `minion_opts`.
1. k8sglue uses `salt-ssh` to:
    1. Install `salt-minion` and configure it to use the salt-master(s) from the cluster config.
1. Wait for one of the salt-master(s) to accept the key.
* Results in: One or more nodes to be prepared and joined in the cluster.

### Automatic machine join to cluster

#### Requirements for a machine

* `salt-minion` installed and configured to use the salt-master(s) of the cluster.
* Must have the Salt SSH deploy public key allowed to connect as `root` or other user that can do passwordless `sudo`.
    * **NOTE**: This is needed so the salt-master(s) can verify the salt-minion's public key.

#### Instructions

1. The machine must have `salt-minion` installed which is already configured to go to one of the salt-master(s).
    1. One can use `k8sglue machine join` or (when written) `salt-ssh` to trigger the configuration.
1. Machine's salt-minion connects to salt-master(s), an event will trigger and cause the salt-master to run `salt-run salt.runners.manage.safe_accept` for that machine.
    1. On success, a high state will be triggered for that machine.
    1. On connection failure, it was probably a "bad" machine that tried to connect to the salt-master(s).
* Results in: a new machine added to the cluster.

## Goals
### Primary Goals
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

## k8sglue Commands and Subcommands

* `salt` - Salt master management command
    * `apply` - Trigger salt high state.
    * `certs` - Generate and sync certs for the salt-master(s) (if needed and forceable by flag).
    * `init` - Init the salt-master(s) by installing and configuring them.
    * `roster` - Print out the generated salt-master(s) roster file by looking at the cluster config file.
    * `sync` - Sync current (given) `salt` directory to all salt-master(s).
    * `ping` - Run `test.ping` using `salt-ssh`.
* `machines` - Machines management command
    * `join` - Prepare and join one or more nodes to use the salt-master(s).
* `completion`
    * `bash` - Output BASH completion
    * `zsh` - Output ZSH completion
* `help` - Show help menu.
