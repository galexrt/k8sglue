{%- set minion_to_join = salt['pillar.get']('minion_to_join') %}
{%- set kubernetes_nodes = [] %}
{%- set kubernetes_nodes_list = salt['mine.get']('roles:kubernetes-master-init', 'kubernetes_nodes', tgt_type='grain').values() %}
{%- if kubernetes_nodes_list|length != 0 %}
{%- set kubernetes_nodes = kubernetes_nodes_list.join('\n').split('\n') %}
{%- endif %}

{%- if minion_to_join not in kubernetes_nodes %}
send custom kubernetes want-to-join event:
  salt.runner:
    - name: event.send
    - tag: custom/kubernetes/want-to-join
    - data:
        minion_to_join: '{{ minion_to_join }}'
{%- endif %}
