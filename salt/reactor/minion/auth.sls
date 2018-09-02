# TODO Add simpe "ssh to new machines to verify and add the public key to all masters" mechanism
# 1. Simple single state to run ssh from one of the salt-master machines.
#      * Don't use saltify for nodes that already have Salt minion and have connected to a master. Go to Step 2 for that.
# 2. Use orchestrated run(s) to distribute to the machines
#
# Conceptual Plan for "automating" adding "new" nodes:
# * salt-cloud saltify should either be used manually by the user or "automatically" called through the salt-api.
# * If a new node connects to one of the salt-master servers, the following steps will be executed (partly see above as a guideline too):
#     * Connect to the host using `salt-ssh` or "plain" ssh to get the minion's public key.
#     * After the minion's public key has been retrieved an orchestrated run is started to "push" the key to all salt-master servers.
{% if 'act' in data and data['act'] == 'pend' %}
call salt-master minion key verify orch:
  runner.state.orchestrate:
    - args:
      - mods: orch.salt-key
      - pillar:
          minion: {{ data['id'] }}
          pub_key: {{ data['pub'] }}
{% endif %}
