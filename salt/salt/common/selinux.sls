{% if salt['grains.get']('os_family', 'N/A') == 'RedHat' %}
{% set selinux = salt['pillar.get']('selinux', {}) -%}
selinux:
  pkg.installed:
    - pkgs:
      - policycoreutils-python
{%- if salt['grains.get']('osmajorrelease', '0') == '7' %}
      - policycoreutils-devel
{%- endif %}

selinux-config:
  file.managed:
    - name: /etc/selinux/config
    - user: root
    - group: root
    - mode: '0600'
    - source: salt://common/templates/etc/selinux/config
    - template: jinja

selinux-state:
    cmd.run:
      - name: setenforce {{ selinux.state|default('enforcing') }}
      - unless: if [ "$(sestatus | awk '/Current mode/ { print $3 }')" = enforcing ]; then /bin/true; else /bin/false; fi
{% endif %}
