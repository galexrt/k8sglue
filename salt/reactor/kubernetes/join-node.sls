{% set roles = salt['grains.get']('roles', []) -%}
{% if "kubernetes-master" in roles or "kubernetes-worker" in roles %}
kubernetes kubeadm join node {{ data['id'] }}:
  local.state.apply:
    - tgt: '{{ data['id'] }}'
    - tgt_type: list
    - args:
      - mods: kubernetes-kubeadm.join
      - pillar:
          new_node: {{ data['id'] }}
{% endif %}
