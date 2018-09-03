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

scp key to master:
  cmd.run:
    - name: |
        scp \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion_to_check }}":/etc/salt/pki/minion/minion.pub \
            "/etc/salt/pki/master/minions/{{ minion_to_check }}"
    - require:
      - cmd: 'check if minion reachable'

remove key from master:
  file.absent:
    - name: "/etc/salt/pki/master/minions_pre/{{ minion_to_check }}"
    - require:
      - cmd: scp key to master
