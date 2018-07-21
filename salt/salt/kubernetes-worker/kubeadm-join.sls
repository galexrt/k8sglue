enable kubelet service:
  service.enabled:
    - require:
      - pkg: kubelet
    - enable: True

run command:
  cmd.run:
    - name: echo kubeadm join worker > /opt/test
