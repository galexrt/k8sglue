{%- set roles = salt['grains.get']('roles') %}
include:
{%- if roles is not none %}
{%- if "kubernetes_master" in roles or "kubernetes_master_init" in roles %}
    - kubeadm.master
{%- endif %}
{%- if "kubernetes_master_init" in roles %}
    - kubeadm.master_init
{%- endif %}
{%- endif %}
    - kubeadm.worker
