include:
- kubernetes-kubeadm.kubelet-service

send kubeadm token created event:
  event.send:
    - name: custom/kubernetes/kubeadm/token-generated
    - require:
      - service: kubelet
    - data:
        token: {{ salt['cmd.script'](name='salt://kubernetes-kubeadm/templates/scripts/kubeadm-token.sh', template='jinja') }}