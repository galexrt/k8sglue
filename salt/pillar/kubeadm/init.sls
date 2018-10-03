{%- set roles = salt['grains.get']('roles') %}
{%- if roles is not none %}
include:
{%-     if "kubernetes_master" in roles %}
    - kubeadm.master
{%-     endif %}
    []
{%- endif %}
