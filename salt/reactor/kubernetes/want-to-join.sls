{%- set node_to_join = salt['pillar.get']('node') %}
include:
- kubernetes-kubeadm
- kubernetes-kubeadm.kubelet-service

{%- set kubernetes_nodes = salt['mine.get']('roles:kubernetes-master-init', 'kubernetes_nodes', tgt_type='grain').values() %}
{%- if node_to_join in kubernetes_nodes %}
send custom_kubernetes_want-to-join event:
  event.send:
    - name: custom/kubernetes/want-to-join
    - require:
      - service: kubelet
    - data:
        node: '{{ node_to_join }}'
{%- endif %}
