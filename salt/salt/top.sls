{%- set roles = salt['grains.get']('roles') %}
{%- set containerRuntime = salt['pillar.get']('clusterConfig:containerRuntime', "crio") -%}
base:
  '*':
    - common
    - system-mainteance
    - sysctl
    - wireguard
{%- if roles is none or "salt_master" not in roles %}
    - salt-minion
{%- endif %}
  'G@roles:salt_master':
    - salt-master
    - salt-cloud
  'G@roles:kubernetes_*':
    - docker
{%- if containerRuntime == "crio" %}
    - crio
{%- endif %}
    - kubeadm
