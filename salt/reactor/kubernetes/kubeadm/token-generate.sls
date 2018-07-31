{%- set hosts = salt['saltutil.runner']('manage.up', tgt='roles:kubernetes-master' tgt_type='grain', []) %}
{%- set kubernetes_master = hosts|random %}

kubernetes kubeadm generate token for {{ data['id'] }}:
  local.state.single:
    - tgt: '{{ kubernetes_master }}'
    - tgt_type: list
    - args:
      - fun: pkg.installed
      - name: zsh
