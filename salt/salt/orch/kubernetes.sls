{%- set hosts = salt['saltutil.runner']('manage.up', []) %}
{%- set data = salt['pillar.get']('event_data', None) %}
{%- if data != None %}
{{ hosts.append(data.id) }}
{%- endif %}

setup first kubernetes master:
  salt.state:
    - tgt: 'roles:kubernetes-master-init'
    - tgt_type: grain
    - sls:
      - kubernetes-kubeadm.kubeadm

{% for host in hosts %}
get kubeadm token for host {{ host }}:
  salt.state:
    - tgt: 'roles:kubernetes-master-init'
    - tgt_type: grain
    - sls:
      - kubernetes-kubeadm.kubeadm-token
    - pillar:
      hosts: {{ host }}

setup kubernetes worker on {{ host }}:
  salt.state:
    - tgt: 'roles:kubernetes-worker'
    - tgt_type: grain
    - require:
      - salt: 'setup first kubernetes master'
      - salt: 'get kubeadm token for host {{ host }}'
    - sls:
      - kubernetes-kubeadm.kubeadm
{% endfor %}
