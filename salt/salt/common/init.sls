include:
  - common.packages
  - common.dns
  - common.selinux
{%- if grains['swap_total']|default(0)|int > 0 %}
  - common.swap
{%- endif %}
  - common.kmodules
