install salt-master package:
  pkg.latest:
    - name: salt-master
    - refresh: True

install salt-ssh package:
  pkg.latest:
    - name: salt-ssh
    - refresh: True

# TODO use replace to simply change `master` and `ssl` config in master config file
# no matter what OS and/or changes happen to this file
configure salt-master:
  file.managed:
    - name: /etc/salt/master
    - user: root
    - group: root
    - mode: 644
    - template: jinja
    - source: salt://salt-master/master

start salt-master:
  service.running:
    - name: salt-master
    - require:
      - pkg: salt-master
    - watch:
      - file: /etc/salt/master
    - enable: True
