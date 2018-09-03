{% set minion = salt['pillar.get']('minion') %}
run minion ssh key verify:
  salt.state:
    - tgt: 'roles:salt-master'
    - tgt_type: grain
    - sls:
      - glue.minion_key_verify
    - pillar:
      - minion: {{ minion }}
