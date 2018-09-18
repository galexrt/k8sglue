include:
- kubeadm
- kubeadm.kubelet-service

kubeadm init:
  cmd.script:
    - source: salt://kubeadm/templates/scripts/kubeadm-init.sh
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
