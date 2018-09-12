{%- set kubernetes_master = salt['saltutil.runner']('manage.up', ['roles:kubernetes-master-init', 'grain'])|random %}
{%- set node = data['id'] %}

'kubeadm generate token for {{ node }}':
  local.state.apply:
    - tgt: '{{ kubernetes_master }}'
    - tgt_type: list
    - sync_mods: all
    - args:
      - mods: kubernetes-kubeadm.token-generate
      - pillar:
          node: {{ node }}
