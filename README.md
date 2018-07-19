# k8sglue
This project is a total work in progress right now!

## Requirements

* Saltstack (`salt-ssh`)
* Golang to compile the project (only tested with >= `1.10`)

## k8sglue Commands and Subcommands

* `salt` - Salt master management command
    * `apply` - Trigger salt high state.
    * `certs` - Generate and sync certs for the salt-master(s) (if needed and forceable by flag).
    * `init` - Init the salt-master(s) by installing and configuring them.
    * `keys` - Salt keys management command.
        * `accept` - Accept the salt-key of one or more machines on all salt-master(s).
        * `remove` - Remove the salt-key of one or more machines on all salt-master(s).
    * `ping` - Run `test.ping` using `salt-ssh`.
    * `roster` - Print out the generated salt-master(s) roster file by looking at the cluster config file.
    * `sync` - Sync current (given) `salt` directory to all salt-master(s).
* `machines` - Machines management command
    * `prepare` - Prepare one or more nodes by using salt-ssh to run the `base` states (`common` and `salt-minion`). In the end the node's salt-minion must be connected to the salt-master(s).
* `completion`
    * `bash` - Output BASH completion
    * `zsh` - Output ZSH completion
* `help` - Show help menu.

#### Status

| Command                    | Status |
| -------------------------- | ------ |
| `k8sglue salt apply`       | Works² |
| `k8sglue salt certs`       | TODO   |
| `k8sglue salt init`        | WIP    |
| `k8sglue salt keys`        | Done¹  |
| `k8sglue salt keys accept` | Works² | 
| `k8sglue salt keys remove` | WIP    |
| `k8sglue salt ping`        | Works² |
| `k8sglue salt roster`      | Works² |
| `k8sglue salt sync`        | WIP    |
| `k8sglue machines prepare` | WIP    |
| `k8sglue completion bash`  | Done¹  |
| `k8sglue completion zsh`   | Done¹  |
| `k8sglue help`             | Done¹  |

¹ The function of the command (if any) is unlikely to change it's behavior and is not a "core" component.
² Command is currently "stable" but it's behavior may change in the future.

### `k8sglue salt apply`

1. k8sglue uses `salt-ssh` to trigger the Salt highstate from one of the salt-master(s).
* Results in: One of the salt-master(s) having triggered the  Salt highstate on all nodes.

### `k8sglue salt certs`

1. k8sglue generate SSL certificates with all the DNS names and IPs for the given machines that should be salt-master(s).
1. k8sglue synces SSL certificates for the salt-master(s) using `salt-ssh`.
* Results in: salt-master(s) certificates generated and synced to the salt-master(s).

> **NOTE** This needs to be run when new salt-master(s) are added.

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
1. k8sglue uses a `salt key accept` mechanism to get all minion fingerprints and accept them on each salt-master(s).
    1. See `k8sglue machine join`.
* Results in: salt-master(s) ready for use and "connected to themselves".

> **NOTE** Each new node must connect to the salt-master(s).
> Each node will be "tested" with `salt.runners.manage.safe_accept` and accepted if the fingerprints match.

### `k8sglue salt keys accept`

1. ks8glue uses `salt-ssh` to get the salt-minion's fingerprint hash of each given machine.
1. k8sglue uses `salt-ssh` to then accept the machine by name (should be improved in the future).
    1. Connect to each salt-master(s) and run a simple bash script to check with `salt-key -f MACHINE` if the key is already accepted.
    1. If the key is not accepted, run `salt-key` to accept it.
* Results in: The given machine(s)'s salt-minion key(s) are accepted on all salt-master(s).

### `k8sglue salt keys remove`

1. k8sglue uses `salt-ssh` to then accept the machine by name (should be improved in the future).
    1. Connect to each salt-master(s) and simply run `salt-key -q -y -d MACHINE` to delete the key if it exists.
* Results in: The given machine(s)'s salt-minion key(s) are removed on all salt-master(s).

### `k8sglue salt ping`

1. k8sglue runs `salt-ssh` on all machines a given pattern matches (takes one flag which is the name pattern given to `salt-ssh`).
* Results in: Salt `test.ping` being run on every machine

### `k8sglue salt roster`

1. k8sglue generate salt-ssh roster file.
* Results in: An usable salt roster file which can be used by anyone to connect to the machines as wanted.

### `k8sglue salt sync`

1. k8sglue uses `salt-ssh state.single file.recurse` to copy the directory to each salt-master(s).
* Results in: Sync the given Salt states directory to each salt-master(s).

> **NOTE** Look into if it may also be a good thing to sync the pillars and mines between the salt-master(s).

### `k8sglue machine prepare`

1. k8sglue writes a Saltstack Roster file with given roles for the node(s) given specified as `minion_opts`, but only triggering salt-minion and common state.
1. k8sglue uses `salt-ssh` to:
    1. Run the base install ( which includes the `common` and `salt-minion` states).
    * Results in: Configured salt-minion and basic configured machine.
1. k8sglue uses `salt-ssh` to get the salt-minion public key.
1. Connect to each salt-master(s) and accept the public key.
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
