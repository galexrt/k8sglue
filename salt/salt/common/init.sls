include:
  - common.timezone
  - common.dns
  - common.selinux
{%- if salt['grains.get']('swap_total', '0')|int > 0 %}
  - common.swap
{%- endif %}
  - common.kmodules

install common packages:
  pkg.installed:
    - pkgs: {{ salt['pillar.get']('common:packages', []) }}
