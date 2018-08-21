include:
- kubernetes-kubeadm
- kubernetes-kubeadm.kubelet-service

send kubeadm token created event:
  event.send:
    - name: custom/kubernetes/token-generated
    - require:
      - service: kubelet
    - data:
        token: '{{ salt['cmd.script']('salt://kubernetes-kubeadm/templates/scripts/kubeadm-token.sh', template='jinja')['stdout'] }}'
        node: '{{ salt['pillar.get']('node') }}'
