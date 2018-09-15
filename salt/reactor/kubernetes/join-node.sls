{%- set minion_to_join = data['data']['minion_to_join'] %}
{%- set token = data['data']['token'] %}
kubeadm join node {{ node }}:
  local.state.apply:
    - tgt: '{{ minion_to_join }}'
    - tgt_type: list
    - sync_mods: all
    - args:
      - mods: kubernetes-kubeadm.join
      - pillar:
          minion_to_join: '{{ minion_to_join }}'
          token: '{{ token }}'
