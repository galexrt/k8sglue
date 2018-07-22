enable kubelet service:
  service.enabled:
    - name: kubelet
    - enable: True

kubeadm init:
  cmd.script:
    - source: salt://kubernetes-master/scripts/kubeadm-init.sh
    - template: jinja
    - creates: /etc/kubernetes/manifests/kube-apiserver.yaml
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
