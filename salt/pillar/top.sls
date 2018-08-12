---
{% set roles = salt['grains.get']('roles') -%}
base:
  '*':
    - headers
    - common
    - network
{% if roles is not none %}
{% if "kubernetes-master-init" in roles or "kubernetes-master" in roles or "kubernetes-worker" in roles %}
    - kubernetes-kubeadm
{% endif %}
{% endif %}
