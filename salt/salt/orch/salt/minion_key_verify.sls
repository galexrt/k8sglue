{%- set minion_to_check = salt['pillar.get']('minion_to_check') %}
{%- set random_master = salt['manage.up'](tgt="roles:salt_master", tgt_type="grain")|random %}

'check if minion {{ minion_to_check }} is reachable by ssh':
  salt.state:
    - tgt: 'roles:salt_master'
    - tgt_type: grain
    - sls:
      - k8sglue.salt.minion.check_minion_ssh_reachability
    - pillar:
        minion_to_check: {{ minion_to_check }}

'stop minion {{ minion_to_check}}':
  salt.state:
    - tgt: '{{ random_master }}'
    - tgt_type: list
    - sls:
      - k8sglue.salt.minion.stop_salt_minion
    - require:
      - salt: 'check if minion {{ minion_to_check }} is reachable by ssh'

'copy minion {{ minion_to_check }} ssh key':
  salt.state:
    - tgt: 'roles:salt_master'
    - tgt_type: grain
    - sls:
      - k8sglue.salt.minion.copy_minion_key
    - pillar:
        minion_to_check: {{ minion_to_check }}
    - require:
      - salt: 'stop minion {{ minion_to_check}}'

'start minion {{ minion_to_check}} again':
  salt.state:
    - tgt: '{{ random_master }}'
    - tgt_type: list
    - sls:
      - k8sglue.salt.minion.restart_salt_minion
    - require:
      - salt: 'copy minion {{ minion_to_check }} ssh key'
