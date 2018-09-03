{% set minion = salt['pillar.get']('minion') %}
check if minion reachable:
  cmd.run:
    - name: |
        ssh \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion }}" \
            'echo "I IS REACHABLE"'

scp key to master:
  cmd.run:
    - name: |
        scp \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion }}":/etc/salt/pki/minion/minion.pub \
            "/etc/salt/pki/master/minions/{{ minion }}"
    - require:
      - cmd: 'check if minion reachable'

remove key from master:
  file.absent:
    - name: "/etc/salt/pki/master/minions_pre/{{ minion }}"
    - require:
      - cmd: scp key to master
