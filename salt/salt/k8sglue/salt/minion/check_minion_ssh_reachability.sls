{%- set minion_to_check = salt['pillar.get']('minion_to_check') %}
check if minion reachable:
  cmd.run:
    - name: |
        ssh \
            -o UserKnownHostsFile=/dev/null \
            -o StrictHostKeyChecking=no \
            -i /etc/salt/ssh/id_rsa \
            "{{ minion_to_check }}" \
            'echo "I IS REACHABLE"'
