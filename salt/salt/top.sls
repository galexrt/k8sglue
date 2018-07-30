{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
base:
  '*':
    - common
    - sysctl
    - system-maintenance
    - salt-minion
  'G@roles:salt-master':
    - salt-master
  'G@roles:kubernetes-*':
    - kubernetes-kubeadm.basics
{%- if containerRuntime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
