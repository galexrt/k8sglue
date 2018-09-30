{%- set roles = salt['grains.get']('roles') %}
{%- set containerRuntime = salt['pillar.get']('cluster_config:containerRuntime', "crio") -%}
base:
  '*':
    - common
    - sysctl
    - system-maintenance
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
#  'G@roles:kubernetes_master_init':
#    - kubeadm.masterinit
