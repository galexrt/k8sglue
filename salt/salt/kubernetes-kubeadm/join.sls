{%- set roles = salt['grains.get']('roles', []) %}
{%- if "kubernetes-master-init" not in roles %}
include:
- kubernetes-kubeadm
- kubernetes-kubeadm.kubelet-service

kubeadm join:
  cmd.script:
    - source: salt://kubernetes-kubeadm/templates/scripts/kubeadm-join.sh
    - template: jinja
    - env:
      - KUBEADM_JOIN_TOKEN: '{{ salt['pillar.get']('token') }}'
    - creates: /etc/kubernetes/pki/ca.crt
    - require:
      - service: kubelet

{%- if "kubernetes-master" in roles %}
copy kubeconfig to /root/.kube:
  file.copy:
    - name: /root/.kube/config
    - source: /etc/kubernetes/admin.conf
    - makedirs: True
    - user: root
    - group: root
    - require:
      - cmd: 'kubeadm join'
{%- endif %}
{%- endif %}

ROLES {{ roles }}
TOKEN {{ salt['pillar.get']('token') }}
