{%- set roles = salt['grains.get']('roles', []) %}
include:
- kubeadm
- kubeadm.kubelet-service

kubeadm join:
  cmd.script:
    - source: salt://kubeadm/templates/scripts/kubeadm-join.sh
    - template: jinja
    - env:
      - KUBEADM_JOIN_TOKEN: '{{ salt['pillar.get']('token') }}'
    - creates: /etc/kubernetes/pki/ca.crt
    - require:
      - service: kubelet

{%- if "kubernetes_master" in roles %}
copy kubeconfig to /root/.kube:
  file.copy:
    - name: /root/.kube/config
    - source: /etc/kubernetes/admin.conf
    - makedirs: true
    - user: root
    - group: root
    - require:
      - cmd: 'kubeadm join'

push kubernetes ca cert to master:
  module.run:
    - name: cp.push
    - path: /etc/kubernetes/pki/ca.crt
    - require:
      - cmd: 'kubeadm join'
{%- endif %}
