{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
enable kubelet service:
  service.enabled:
    - name: kubelet
    - enable: True

kubeadm join:
  cmd.run:
    - name: kubeadm join --feature-gates=DynamicKubeletConfig {% if containerRuntime == "crio" %} --cri-socket=/var/run/crio/crio.sock{% endif %} > /opt/test
    - creates: /etc/kubernetes/manifests/kube-apiserver.yaml
