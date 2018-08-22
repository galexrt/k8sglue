{%- set roles = salt['grains.get']('roles') %}
include:{% if not ("kubernetes-master-init" in roles or "kubernetes-master" in roles or "kubernetes-worker" in roles) %} []{% endif %}
{% if roles is not none %}
{% if "kubernetes-master" in roles %}
    - kubernetes-kubeadm.master
{% endif %}
{% if "kubernetes-master-init" in roles %}
    - kubernetes-kubeadm.master-init
{% endif %}
{% endif %}
