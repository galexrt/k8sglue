{% set minion_to_check = salt['pillar.get']('minion_to_check') %}
copy master sign pub key to minion:
  salt.state:
    - tgt: 'roles:salt-master'
    - tgt_type: grain
    - sls:
      - glue.master_pub_key
    - pillar:
        minion_to_check: {{ minion_to_check }}

run minion ssh key verify:
  salt.state:
    - tgt: 'roles:salt-master'
    - tgt_type: grain
    - sls:
      - glue.minion_key_verify
    - pillar:
        minion_to_check: {{ minion_to_check }}
    - require:
      - state: 'copy master sign pub key to minion'
