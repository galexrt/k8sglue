# TODO Add simpe "ssh to new machines to verify and add the public key to all masters" mechanism
# 1. Simple single state to ssh run from one of the salt-master machines.
#      * Don't use saltify for nodes that already have Salt minion and have connected to a master. Go to Step 2 for that.
# 2. Use orchestrated run(s) to distribute to the machines
