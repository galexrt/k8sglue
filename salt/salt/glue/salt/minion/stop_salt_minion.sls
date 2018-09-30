{%- set minion_to_check = salt['pillar.get']('minion_to_check') %}
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
