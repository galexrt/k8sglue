{% set roles = salt['grains.get']('roles', []) -%}
{% if "kubernetes-master" in roles or "kubernetes-worker" in roles %}
include:
- kubernetes-kubeadm.kubelet-service

kubeadm join:
  cmd.script:
    - source: salt://kubernetes-kubeadm/templates/scripts/kubeadm-join.sh
    - template: jinja
    - creates: /etc/kubernetes/pki/ca.crt
    - require:
      - service: kubelet

{% if "kubernetes-master" in roles %}
copy kubeconfig to /root/.kube:
  file.copy:
    - name: /root/.kube/config
    - source: /etc/kubernetes/admin.conf
    - makedirs: True
    - user: root
    - group: root
    - require:
      - cmd: 'kubeadm join'
{% endif %}
{% endif %}
