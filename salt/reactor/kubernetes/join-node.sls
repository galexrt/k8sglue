{%- set node = data['data']['node'] %}
{%- set token = data['data']['token'] %}
kubeadm join node {{ node }}:
  local.state.apply:
    - tgt: '{{ node }}'
    - tgt_type: list
    - sync_mods: all
    - args:
      - mods: kubernetes-kubeadm.join
      - pillar:
          node: '{{ node }}'
          token: '{{ token }}'
