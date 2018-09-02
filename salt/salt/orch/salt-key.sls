minion get key by ssh:
  cmd.run:
    - tgt: 'roles:salt-master'
    - tgt_type: grain
    - args:
      - cmd: |
          ssh \
              -o UserKnownHostsFile=/dev/null \
              -o StrictHostKeyChecking=no \
              "{{ data['id'] }}" \
              'echo I IS REACHABLE'
          KEY="$(ssh \
              -o UserKnownHostsFile=/dev/null \
              -o StrictHostKeyChecking=no \
              "{{ data['id'] }}" \
              'systemctl start salt-minion && salt-call key.finger --local --out txt --no-color | sed -n -e "s/^local: //p"')"
          echo "Key fingerprint is ${KEY}"
