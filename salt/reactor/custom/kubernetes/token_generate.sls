{%- set kubernetes_master = salt['mine.get']('roles:kubernetes_master', 'ipv4_addresses', tgt_type='grain').values() %}
{%- set minion_to_join = data['minion_to_join'] %}

TEST {{ kubernetes_master }}
{{ salt['mine.get']('roles:kubernetes_master', 'ipv4_addresses', tgt_type='grain') }}

'kubeadm generate token for {{ minion_to_join }}':
  local.state.apply:
    - tgt: '{{ kubernetes_master }}'
    - tgt_type: list
    - sync_mods: all
    - args:
      - mods: kubeadm.token-generate
      - pillar:
          minion_to_join: '{{ minion_to_join }}'
