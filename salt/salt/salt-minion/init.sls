salt-minion python additional feature dependencies:
  pkg.latest:
    - pkgs:
      - python-croniter
      - python-dateutil
    - refresh: true

install salt-minion package:
  pkg.latest:
    - name: salt-minion
    - refresh: true

configure salt-minion:
  file.recurse:
    - name: /etc/salt/minion.d
    - source: salt://salt-minion/templates/etc/salt/minion.d
    - user: root
    - group: root
    - dir_mode: 640
    - file_mode: 750
    - replace: true
    - clean: false
    - template: jinja

start salt-minion:
  service.running:
    - name: salt-minion
    - require:
      - pkg: salt-minion
      - pkg: 'salt-minion python additional feature dependencies'
    - watch:
      - file: 'configure salt-minion'
    - enable: true
