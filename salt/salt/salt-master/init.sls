install salt-master package:
  pkg.latest:
    - name: salt-master
    - refresh: True

install salt-ssh package:
  pkg.latest:
    - name: salt-ssh
    - refresh: True

configure salt-master:
  file.managed:
    - name: /etc/salt/master
    - user: root
    - group: root
    - mode: 644
    - template: jinja
    - source: salt://salt-master/etc/salt/master

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

start salt-master:
  service.running:
    - name: salt-master
    - require:
      - pkg: salt-master
    - watch:
      - file: /etc/salt/master
    - enable: True
