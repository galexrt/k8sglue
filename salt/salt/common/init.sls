include:
  - common.packages
  - common.dns
  - common.selinux
{%- if salt['grains.get']('swap_total', '0')|int > 0 %}
  - common.swap
{%- endif %}
  - common.kmodules

install common packages:
  pkg.installed:
    - pkgs:
      - python3-dnf-plugin-tracer
      - htop
      - iftop
      - iotop
      - sysstat
      - tcpdump
      - conntrack-tools
      - ipvsadm
