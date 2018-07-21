{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
enable kubelet service:
  service.enabled:
    - name: kubelet
    - enable: True

kubeadm join:
  cmd.run:
    - name: echo kubeadm join {% if containerRuntime == "crio" %} --cri-socket=/var/run/crio/crio.sock{% endif %} > /opt/test
    - unless: test -f /etc/kubernetes/manifests/kube-apiserver.yaml
