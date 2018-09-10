include:
- salt-minion

{% for package in ['salt-master', 'salt-api', 'salt-ssh'] %}
install {{ package }} package:
  pkg.latest:
    - name: {{ package }}
    - refresh: True
{% endfor %}

configure salt-master:
  file.recurse:
    - name: /etc/salt/master.d
    - source: salt://salt-master/templates/etc/salt/master.d
    - user: root
    - group: root
    - dir_mode: 640
    - file_mode: 750
    - replace: True
    - clean: True
    - template: jinja

{% for dir in ['/etc/salt/roster.d', '/etc/salt/ssh'] %}
create {{ dir }} directory for salt-master:
  file.directory:
    - name: {{ dir }}
    - user: root
    - group: root
    - dir_mode: 750
    - file_mode: 640
    - recurse:
      - user
      - group
      - mode
    - require:
      - pkg: salt-master
{% endfor %}

generate master signing signature:
  cmd.run:
    - name: |
        set -e
        salt-key --auto-create --gen-keys=master --gen-keys-dir=/etc/salt/pki/master/
        salt-key --gen-signature --auto-create
    - creates:
      - /etc/salt/pki/master/master.pem
      - /etc/salt/pki/master/master.pub
      - /etc/salt/pki/master/master_pubkey_signature
      - /etc/salt/pki/master/master_sign.pem
      - /etc/salt/pki/master/master_sign.pub
    - require:
      - pkg: salt-master
    - require_in:
      - service: 'start salt-master'

generate salt-master minion keys:
  cmd.run:
    - name: |
        set -e
        salt-key --auto-create --gen-keys=minion --gen-keys-dir=/etc/salt/pki/minion/
    - creates:
      - /etc/salt/pki/minion/minion.pem
      - /etc/salt/pki/minion/minion.pub
    - require:
      - pkg: salt-master
      - pkg: salt-minion
    - require_in:
      - service: 'start salt-master'
      - service: 'start salt-minion'

copy master minion key to accepted:
  file.symlink:
    - name: '/etc/salt/pki/master/minions/{{ salt['grains.get']('fqdn') }}'
    - target: /etc/salt/pki/minion/minion.pub
    - makedirs: True
    - require:
      - pkg: salt-master
      - pkg: salt-minion
    - require_in:
      - service: 'start salt-minion'
      - service: 'start salt-master'

start salt-master:
  service.running:
    - name: salt-master
    - require:
      - pkg: salt-master
      - cmd: 'generate master signing signature'
    - watch:
      - file: 'configure salt-master'
    - enable: True

start salt-api:
  service.running:
    - name: salt-api
    - require:
      - pkg: salt-api
      - service: 'start salt-master'
    - watch:
      - file: 'configure salt-master'
    - enable: True
