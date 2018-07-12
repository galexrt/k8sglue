# K8SGLUE v1
This project is work in progress!

This file currently mostly contains the concept/idea/flows behind this project.

## Flows
### Initial Cluster Creation Flow
#### Getting machines from Terraform

> **NOTE** This must be run only once!

1. Terraform has been run and provisioned machines (only a salty touchy the machines, only `dnf update -y` and `dnf install salt-minion -y`! (if possible otherwise will be triggered by salt state `salt-minion`)).
1. Terraform provides outputs that can be used by the GLUE.
1. GLUE subcommand to take Terraform servers output and "transform" it into a node list.

* Results in: Node list which can be consumed by GLUE.

#### Creating and initiating the components for the Cluster

1. GLUE builds list of machines, get machines with `roles.salt.master: true` for the next step.
1. GLUE checks if salt-master node(s) already have a certificate, if so download.
    1. GLUE creates or renews salt-master(s) certificates.
1. GLUE uses `salt-ssh` to trigger the `salt-master` state on each Kubernetes-master and salt-master node(s).
    1. For the begining it should be enough to have the salt-master(s) run on the same node as the Kubernetes masters.
    1. Create a file roster with `minion_opts` for `salt-ssh`, so that the roles grain can be used and set in the `salt-minion` config.
    1. Copy `salt` directory to `salt-master` node(s) `/srv/salt`.
    1. Wait for `salt-master`s to be ready.

* Result in: `salt-master`s and `kubernetes-master` ready for use.

### Adding a node(s) to a Cluster

1. GLUE generate kubeadm join token (+ certificate).
    1. Either the `kubeadm init` output is used or the certificate is generated locally and then copied to all Kubernetes masters (`openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'`).
        * See https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/#custom-certificates.
1. GLUE uses `salt-ssh` on all machines to apply `salt-minion` state which configures the `salt-minion`.
    1. Use `salt-ssh` roster `minion_opts` to set the roles grain.
    * Results in: all machines having their roles set as a static minion grain and are connecting to one of the salt-master(s).

1. GLUE uses `salt-run salt.runners.manage.safe_accept` to accept hosts (a (second) private ssh key for that is needed).
    1. GLUE runs `salt-ssh` on all machines to create a symlink from `/etc/salt/pki/minion` to `/var/tmp/.salt/running_data/etc/salt/pki/minion`. Until https://github.com/saltstack/salt/issues/37474 is fixed.
    1. GLUE then uses `salt-run salt.runners.manage.safe_accept` (https://docs.saltstack.com/en/latest/ref/runners/all/salt.runners.manage.html) or try to use it directly in a state file which is only used for the salt-master(s).

1. GLUE Run salt `ping` on all machines to verify they are reachable.
1. GLUE Salt high state is triggered or all new machines, which configures all new machines with current state.
    1. Satl magic happens, so that kubeadm and other magic is installed on the servers as needed given by their role grain(s).

1. Salt high state is triggered in an interval of 30 minutes to ensure everything is setup as wanted.

* Results in: New and existing machines to be (re-)configured.

## Goals
### Secondary Goals
* GLUE has `machines` subcommand for adding, resetting and removing a machine (would trigger key removal in Salt and remove the machine from the Kubernetes cluster).
    * Use `salt-ssh` to do the job.
* Salt Beacons should be implemented for additional monitoring and reacting: "automatic problem solver" (e.g. a node runs out of memory = run `logrotate`, a node is not ready with kubelet PLEG errors = reboot node).
* Servers completely (host) firewalled (with exceptions, e.g. role loadbalancer which allows certain ports).

## What is used for what
* Terraform
    * Provisions "cloud" machines (everything VM, server, whatever that can be provisioned)
    * Creation and management of certificates
* GLUE
    * Creates certificates for salt-master(s).
    * Creates kubeadm token (default TTL 4h).
    * Copies Salt files to the salt-master(s).
    * Triggers Salt.

### GLUE Commands and Subcommands
* `cluster` - Cluster management command
    * `init` - Init a cluster by deploying salt-master(s)
    * `salt` - Trigger salt on the salt-master(s)
    * `status` - Show status of a cluster
    * `delete` - Delete a cluster (undeploy salt-master(s))
* `completion`
    * `bash` - Output BASH completion
    * `zsh` - Output ZSH completion
* `config` - Config management command
    * `create`
* `machines` - Machines management command
    * `add` - Add a machine to a cluster (will take a nodes list yaml or a single node by flags)
    * `delete` - Remove a machine to a cluster
    * `info` - Show info about a machine in the cluster it is in
    * `genlist` - genlist command
        * `terraform` - "Translate" Terraform outputs to a YAML machine list for GLUE
        * Other "inputs" will be available too in the future for "auto generation" of machine lists.
* `help` - Show help menu.
