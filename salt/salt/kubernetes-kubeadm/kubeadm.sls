{% set roles = salt['grains.get']('roles', []) %}
enable kubelet service:
  service.enabled:
    - name: kubelet
    - enable: True

kubeadm:
  cmd.script:
{%- if "kubernetes-master-init" in roles %}
    - source: salt://kubernetes-kubeadm/scripts/kubeadm-init.sh
{%- else %}
    - source: salt://kubernetes-kubeadm/scripts/kubeadm-join.sh
{%- endif %}
    - template: jinja
    - creates: /etc/kubernetes/manifests/kube-apiserver.yaml
    - require:
      - service: kubelet

{%- if "kubernetes-master-init" in roles %}
update salt mine kubernetes-master-ca-cert-hash:
  module.run:
    - name: mine.update
    - require:
      - cmd: 'kubeadm'
{%- endif %}

{%- if not "kubernetes-master-init" in roles and salt['grains.get']('host', None) != None %}
delete kubernetes-kubernetes-join-token in salt-mine:
  module.run:
    - name: mine.delete
    - m_fun: kubernetes-join-token-{{ host }}
{%- endif %}

{% if "kubernetes-master" in roles %}
copy kubeconfig to /root/.kube:
  file.copy:
    - name: /root/.kube/config
    - source: /etc/kubernetes/admin.conf
    - makedirs: True
    - user: root
    - group: root
    - require:
      - cmd: 'kubeadm'
{%- endif %}
