---
{%- set roles = salt['grains.get']('roles') %}
base:
  '*':
    - headers
    - common
    - network
{% if roles is not none %}
{% if "kubernetes-master" in roles %}
    - kubernetes-kubeadm.master
{% endif %}
{% if "kubernetes-master-init" in roles %}
    - kubernetes-kubeadm.master-init
{% endif %}
{% endif %}
