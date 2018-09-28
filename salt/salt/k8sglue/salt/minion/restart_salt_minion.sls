{% set minion_to_check = salt['pillar.get']('minion_to_check') %}
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
