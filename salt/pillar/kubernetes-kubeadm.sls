{% set roles = salt['grains.get']('roles', []) -%}
include:{% if not ("kubernetes-master-init" in roles or "kubernetes-master" in roles) %} []{% endif %}
{%- if "kubernetes-master-init" in roles %}
    - kubernetes-kubeadm.master-init
{%- endif %}
{%- if "kubernetes-master" in roles %}
    - kubernetes-kubeadm.master
{%- endif %}
