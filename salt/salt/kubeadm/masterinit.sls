include:
- kubeadm
- kubeadm.kubelet-service

kubeadm init:
  cmd.script:
    - source: salt://kubeadm/templates/scripts/kubeadm-init.j2
    - template: jinja
    - creates: /etc/kubernetes/pki/ca.crt
    - require:
      - service: kubelet

copy kubeconfig to /root/.kube:
  file.copy:
    - name: /root/.kube/config
    - source: /etc/kubernetes/admin.conf
    - makedirs: true
    - user: root
    - group: root
    - require:
      - cmd: 'kubeadm init'

push kubernetes ca cert to master:
  module.run:
    - name: cp.push
    - path: /etc/kubernetes/pki/ca.crt
    - require:
      - cmd: 'kubeadm init'
