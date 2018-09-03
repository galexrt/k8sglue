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
      - cmd: 'check if minion reachable'
