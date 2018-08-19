{% set roles = salt['grains.get']('roles', []) -%}
{% if "kubernetes-master" in roles or "kubernetes-worker" in roles %}
kubeadm join node {{ data['id'] }}:
  local.state.apply:
    - tgt: '{{ data['id'] }}'
    - tgt_type: list
    - sync_mods: all
    - args:
      - mods: kubernetes-kubeadm.join
      - pillar:
          node: {{ data['id'] }}
{% endif %}
