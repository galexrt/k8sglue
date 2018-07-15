install salt-minion package:
  pkg.latest:
    - name: salt-minion
    - refresh: True

# TODO use replace to simply change `master` (and `ssl`?) config in minion config file
# no matter what OS and/or changes happen to this file
configure salt-minion:
  file.managed:
    - name: /etc/salt/minion
    - user: root
    - group: root
    - mode: 644
    - template: jinja
    - source: salt://salt-minion/etc/salt/minion

start salt-minion:
  service.running:
    - name: salt-minion
    - require:
      - pkg: salt-minion
    - watch:
      - file: /etc/salt/minion
    - enable: True
