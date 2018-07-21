{% if grains['os_family'] == 'RedHat' %}
{% set selinux = pillar.get('selinux', {}) -%}
selinux:
  pkg.installed:
    - pkgs:
      - policycoreutils-python
{%- if grains['osmajorrelease'] == '7' %}
      - policycoreutils-devel
{%- endif %}

selinux-config:
  file.managed:
    - name: /etc/selinux/config
    - user: root
    - group: root
    - mode: 600
    - source: salt://common/etc/selinux/config
    - template: jinja

selinux-state:
    cmd.run:
      - name: setenforce {% if selinux.state == 'disabled' %}0{% else %}{{ selinux.state|default('enforcing') }}{% endif %}
      - unless: if [ "$(sestatus | awk '/Current mode/ { print $3 }')" = {{ selinux.state|default('enforcing') }} ]; then /bin/true; else /bin/false; fi
{% endif %}
