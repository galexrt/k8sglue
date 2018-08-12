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

start salt-master:
  service.running:
    - name: salt-master
    - require:
      - pkg: salt-master
    - watch:
      - file: 'configure salt-master'
    - enable: True

start salt-api:
  service.running:
    - name: salt-api
    - require:
      - pkg: salt-api
    - watch:
      - file: 'configure salt-master'
    - enable: True
