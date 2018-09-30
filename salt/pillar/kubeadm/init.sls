{%- set roles = salt['grains.get']('roles') %}
{%- if roles is not none %}
include:
{%-     if "kubernetes_master" in roles or "kubernetes_master_init" in roles %}
    - kubeadm.master
{%-     endif %}
    []
{%- endif %}
