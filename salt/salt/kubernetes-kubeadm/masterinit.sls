include:
- kubernetes-kubeadm
- kubernetes-kubeadm.kubelet-service

kubeadm init:
  cmd.script:
    - source: salt://kubernetes-kubeadm/templates/scripts/kubeadm-init.sh
    - template: jinja
    - creates: /etc/kubernetes/pki/ca.crt
    - require:
      - service: kubelet

copy kubeconfig to /root/.kube:
  file.copy:
    - name: /root/.kube/config
    - source: /etc/kubernetes/admin.conf
    - makedirs: True
    - user: root
    - group: root
    - require:
      - cmd: 'kubeadm init'

ready to accept kubernetes nodes:
  event.send:
    - name: custom/kubernetes/masterinit-is-ready
    - require:
      - file: 'copy kubeconfig to /root/.kube'
