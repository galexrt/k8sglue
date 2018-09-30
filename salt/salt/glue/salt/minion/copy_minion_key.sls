{%- set minion_to_check = salt['pillar.get']('minion_to_check') %}
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
