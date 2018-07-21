{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
enable kubelet service:
  service.enabled:
    - name: kubelet
    - enable: True

run command:
  cmd.run:
    - name: echo kubeadm join master > /opt/test
