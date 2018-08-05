---
{% set roles = salt['grains.get']('roles', []) -%}
base:
  '*':
    - headers
    - common
    - network
{% if "salt-master" in roles %}
    - salt-master
{% endif %}
{% if "kubernetes-master-init" in roles or "kubernetes-master" in roles or "kubernetes-worker" in roles %}
    - kubernetes-kubeadm
{% endif %}
