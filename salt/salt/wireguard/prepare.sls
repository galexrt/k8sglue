generate wireguard privatekey:
  cmd.run:
    - name: 'wg genkey > /etc/wireguard/privatekey'
    - creates: /etc/wireguard/privatekey
    - cwd: /etc/wireguard
    - umask: '077'
    - require:
      - pkg: 'install wireguard dkms and tools'

generate wireguard publickey:
  cmd.run:
    - name: wg pubkey < /etc/wireguard/privatekey > /etc/wireguard/publickey
    - creates: /etc/wireguard/publickey
    - cwd: /etc/wireguard
    - umask: '077'
    - require:
      - cmd: 'generate wireguard privatekey'

push wireguard public key to salt master:
  module.run:
    - name: cp.push
    - path: /etc/wireguard/publickey
    - upload_path: /wireguard/pub/{{ salt['grains.get']('id') }}
    - remove_source: false
    - require:
      - cmd: 'generate wireguard publickey'
