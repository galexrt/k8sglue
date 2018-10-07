{%- set minion_to_join = salt['pillar.get']('minion_to_join') %}
include:
- kubeadm
- kubeadm.kubelet-service

send kubeadm token created event:
  event.send:
    - name: custom/kubernetes/token-generated
    - require:
      - service: kubelet
    - data:
        token: '{{ salt['cmd.script']('salt://kubeadm/templates/scripts/kubeadm-token.j2', template='jinja')['stdout'] }}'
        minion_to_join: '{{ minion_to_join }}'
