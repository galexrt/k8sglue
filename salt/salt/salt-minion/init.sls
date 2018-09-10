install salt-minion package:
  pkg.latest:
    - name: salt-minion
    - refresh: True

configure salt-minion:
  file.recurse:
    - name: /etc/salt/minion.d
    - source: salt://salt-minion/templates/etc/salt/minion.d
    - user: root
    - group: root
    - dir_mode: 640
    - file_mode: 750
    - replace: True
    - clean: False
    - template: jinja

start salt-minion:
  service.running:
    - name: salt-minion
    - require:
      - pkg: salt-minion
    - watch:
      - file: 'configure salt-minion'
    - enable: True
