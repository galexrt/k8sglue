{%- set node = 'node3' %}
create salt profile:
  file.managed:
    - name: /etc/salt/cloud.profiles.d/salt-{{ node }}.conf
    - source: salt://glue/templates/etc/salt/cloud.profiles.d/profile.conf
    - template: jinja
# Look into https://docs.saltstack.com/en/latest/ref/runners/all/salt.runners.cloud.html
