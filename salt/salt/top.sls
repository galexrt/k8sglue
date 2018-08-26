{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
base:
  '*':
    - common
    - salt-minion
    - sysctl
    - system-maintenance
    - wireguard
  'G@roles:salt-master':
    - salt-master
    - salt-cloud
  'G@roles:kubernetes-*':
{%- if containerRuntime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
    - kubernetes-kubeadm
  'G@roles:kubernetes-master-init':
    - kubernetes-kubeadm.masterinit
