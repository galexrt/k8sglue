{%- set roles = salt['grains.get']('roles') %}
include:
{%- if roles is not none and "kubernetes_master" in roles %}
    - kubeadm.master
{%- else %}
    []
{%- endif %}
