{%- set hosts = salt['saltutil.runner']('manage.up', tgt='roles:kubernetes-master' tgt_type='grain', []) %}
{%- set kubernetes_master = hosts|random %}

kubernetes kubeadm generate token for {{ data['id'] }}:
  local.state.apply:
    - tgt: '{{ kubernetes_master }}'
    - tgt_type: list
    - args:
      - mods: kubernetes-kubeadm.token-generate
      - pillar:
          new_node: {{ data['id'] }}
