kubeadm token create:
  cmd.script:
    - source: salt://kubernetes-kubeadm/templates/scripts/kubeadm-init.sh
    - template: jinja
    - creates: /etc/kubernetes/manifests/kube-apiserver.yaml
    - require:
      - service: kubelet
