{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
enable kubelet service:
  service.enabled:
    - enable: True

run command:
  cmd.run:
    - name: kubeadm init --pod-network-cidr=100.64.0.0/13 --service-cidr=100.72.0.0/16{% if containerRuntime == "crio" %} --cri-socket=/var/run/crio/crio.sock{% endif %}

copy kubeconfig to /root/.kube:
  file.copy:
    - name: /root/.kube/config
    - source: /etc/kubernetes/admin.conf
    - makedirs: True
    - user: root
    - group: root
    - require:
      - cmd: 'run command'
