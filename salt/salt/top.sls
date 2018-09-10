{%- set roles = salt['grains.get']('roles') %}
{%- set containerRuntime = salt['pillar.get']('containerRuntime', "crio") -%}
base:
  '*':
    - common
    - sysctl
    - system-maintenance
    - wireguard
{%- if roles is none or "salt-master" not in roles %}
    - salt-minion
{%- endif %}
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
