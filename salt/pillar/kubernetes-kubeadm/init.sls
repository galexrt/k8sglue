{%- set roles = salt['grains.get']('roles') %}
include:
{%- if roles is not none %}
{%- if "kubernetes-master" in roles or "kubernetes-master-init" in roles %}
    - kubernetes-kubeadm.master
{%- endif %}
{%- if "kubernetes-master-init" in roles %}
    - kubernetes-kubeadm.master-init
{%- endif %}
{%- endif %}
    - kubernetes-kubeadm.worker
