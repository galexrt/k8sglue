{% for package in ['salt-cloud'] %}
install {{ package }} package:
  pkg.latest:
    - name: {{ package }}
    - refresh: true
{% endfor %}

configure salt-cloud config:
  file.recurse:
    - name: /etc/salt/cloud.conf.d
    - source: salt://salt-cloud/templates/etc/salt/cloud.conf.d
    - user: root
    - group: root
    - dir_mode: 640
    - file_mode: 750
    - replace: true
    - clean: true
    - template: jinja

configure salt-cloud providers:
  file.recurse:
    - name: /etc/salt/cloud.providers.d
    - source: salt://salt-cloud/templates/etc/salt/cloud.providers.d
    - user: root
    - group: root
    - dir_mode: 640
    - file_mode: 750
    - replace: true
    - clean: true
    - template: jinja

{% for dir in ['/etc/salt/ssh'] %}
create {{ dir }} directory for salt-cloud:
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
      - pkg: salt-cloud
{% endfor %}
