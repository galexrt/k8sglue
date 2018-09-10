---
base:
  '*':
    - headers
    - common
    - network
{%- if salt['file.file_exists']('/srv/pillar/salt.sls') %}
    - salt
{%- endif %}
    - kubernetes-kubeadm
