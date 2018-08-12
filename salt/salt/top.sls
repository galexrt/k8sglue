{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
base:
  '*':
    - common
    - sysctl
    - system-maintenance
    - salt-minion
  'G@roles:salt-master':
    - salt-master
    - salt-cloud
  'G@roles:kubernetes-*':
    - kubernetes-kubeadm.basics
{%- if containerRuntime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
  'G@roles:kubernetes-master-init':
    - kubernetes-kubeadm.masterinit
