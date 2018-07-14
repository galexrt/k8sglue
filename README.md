# K8SGLUE v1
This project is work in progress!

This file currently mostly contains the concept/idea/flows behind this project.

## Requirements

* Saltstack (`salt-ssh`)

## Flows
### Initial Cluster Creation Flow
#### Getting machines from *any provider*

In this flow, Terraform is used but in the end any tool that can "spit out" a file with the "current" machines (either YAML or JSON) in the "MachineList" format will be fine.

1. Terraform has been run and provisioned machines (only a salty touchy the machines, only `dnf update -y` and `dnf install salt-minion -y`! (if possible otherwise will be triggered by salt state `salt-minion`)).
1. Terraform provides outputs that can be used by the GLUE.
1. GLUE subcommand takes Terraform servers outputs and "transform" them into a machine list `*.yaml`.
* Results in: Node list which can be consumed by GLUE.

#### `k8sglue salt init`

> **NOTE** This should be run only once!

Starting point will be that a "MachineList" is already created containing at least 1 machine with `roles.salt.master: true`.

1. GLUE get lists of machines `CLUSTER_DIR/CLUSTER_NAME/machines/*.yaml`.
    1. Get all machines that have `roles.salt.master: true`.
    1. Generate Saltstack Roster file with `minion_opts` for the salt-master(s) and put it in the tempdir.
    1. Use `salt-ssh '*' test.ping` to check the connectivity. If error, fail.
1. GLUE use `salt-ssh` to get fingerprint, valid names and IPs, and validity time of salt-master(s) certificate.
    1. Try connecting to first salt-master, on failure exit.
        1. In the future allow to specify the certificate and key directly from a host for "disaster recovery".
    1. GLUE keeps, creates or renews the salt-master(s) certificate and key.
1. GLUE generates it's own deploy SSH key (which is used for the connection from salt-master(s) to new machines).
    1. The key is copied to each salt-master(s).
    1. The public key is printed out for the user/admin to put on new machines.
1. GLUE uses `salt-ssh` to trigger the `salt-master` state on the salt-master machine(s).
    1. GLUE makes a copy of the "wanted" `salt` directory.
        1. Add the salt-master(s) certificate and key in it.
        1. Add the deploy SSH key(s) to it.
    1. SFTP the `salt` directory to `salt-master` machine(s) `/srv/salt`.
    1. Wait for all salt-master(s) to be ready.
* Result in: All salt-master(s) ready for use.

> **NOTE** Each new node that connects to the will be "tested" with `salt.runners.manage.safe_accept` and then automatically added to the specific roles it has.

### Adding a machine(s) to the Cluster

#### Requirements for a machine

* Must have the Salt SSH deploy public key allowed to connect as `root`.
* `salt-minion` installed (not configured).

#### Instructions

1. The machine must have `salt-minion` installed which is already configured to go to one of the salt-master(s).
1. Machine's salt-minion connects to salt-master(s), an event will trigger and cause the salt-master to run `salt-run salt.runners.manage.safe_accept` for that machine.
    1. On success, a high state will be triggered for that machine.
    1. On connection failure, it was probably a "bad" machine that tried to connect to the salt-master(s).

## Goals
### Secondary Goals

* GLUE has `machines` subcommand for adding, resetting and removing a machine (would trigger key removal in Salt and remove the machine from the Kubernetes cluster).
    * Use `salt-ssh` to do the job.
* Salt Beacons should be implemented for additional monitoring and reacting: "automatic problem solver" (e.g. a machine runs out of memory = run `logrotate`, a machine is not ready with kubelet PLEG errors = reboot machine).
* Servers are host firewalled (with exceptions, e.g. allow custom which allows certain ports).
    * Or depending on how "good" Canal (Calico) allows to "host firewall" using `GlobalNetworkPolicy`.

## What is used for what

* GLUE
    * Creates certificates for salt-master(s) and kubernetes-master(s).
    * Creates kubeadm token (default TTL 4h).
    * Copies Salt files to the salt-master(s).
    * Triggers Salt.
* Saltstack
    * "Actual" configuration of servers.

### GLUE Commands and Subcommands

* `salt` - Salt management command
    * `init` - Init the salt-master(s) by deploying them
    * `sync` - Sync current `salt` directory to salt-master(s)
    * `certs` - Generate certs for the salt-master(s) (if needed and forceable by flag)
* `completion`
    * `bash` - Output BASH completion
    * `zsh` - Output ZSH completion
* `help` - Show help menu.
