{%- set roles = salt['grains.get']('roles') %}
{%- set containerRuntime = salt['pillar.get']('containerRuntime', "crio") -%}
base:
  '*':
    - common
    - sysctl
    - system-maintenance
    - wireguard-install
{%- if roles is none or "salt_master" not in roles %}
    - salt-minion
{%- endif %}
  'G@roles:salt_master':
    - salt-master
    - salt-cloud
  'G@roles:kubernetes_*':
{%- if containerRuntime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
    - kubeadm
  'G@roles:kubernetes_master_init':
    - kubeadm.masterinit
