{% set minion_to_check = salt['pillar.get']('minion_to_check') %}
check if minion reachable:
  cmd.run:
    - name: |
        ssh \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion_to_check }}" \
            'echo "I IS REACHABLE"'

stop salt-minion on minion:
  cmd.run:
    - name: |
        ssh \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion_to_check }}" \
            'systemctl stop salt-minion'
    - require:
      - cmd: 'check if minion reachable'

scp minion key to master:
  cmd.run:
    - name: |
        scp \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion_to_check }}":/etc/salt/pki/minion/minion.pub \
            "/etc/salt/pki/master/minions/{{ minion_to_check }}"
    - require:
      - cmd: 'stop salt-minion on minion'

scp master sign pub key to minion:
  cmd.run:
    - name: |
        scp \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            /etc/salt/pki/master/master_sign.pub \
            "{{ minion_to_check }}":/etc/salt/pki/minion/master_sign.pub
    - require:
      - cmd: 'stop salt-minion on minion'

remove minion key from master:
  file.absent:
    - name: "/etc/salt/pki/master/minions_pre/{{ minion_to_check }}"
    - require:
      - cmd: 'stop salt-minion on minion'

restart salt-minion on node:
  cmd.run:
    - name: |
        ssh \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion_to_check }}" \
            'systemctl restart salt-minion'
    - require:
      - cmd: 'scp master sign pub key to minion'
      - file: 'remove minion key from master'
