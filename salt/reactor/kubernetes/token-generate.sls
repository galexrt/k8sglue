{%- set hosts = salt['saltutil.runner']('manage.up', ['roles:kubernetes-master', 'grain']) %}
{%- set kubernetes_master = hosts|random %}

kubeadm generate token for {{ data['id'] }}:
  local.state.apply:
    - tgt: '{{ kubernetes_master }}'
    - tgt_type: list
    - args:
      - mods: kubernetes-kubeadm.token-generate
      - pillar:
          node: {{ data['id'] }}
