{% import_yaml "config.yml" as config %}
base:
  '*':
    - common
    - salt-minion
  'roles:kubernetes':
    - kubernetes-basics
{%- if config.container_runtime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
  'roles:salt-master':
    - salt-master
  'roles:kubernetes-master':
    - kubernetes-master
  'roles:kubernetes-worker':
    - kubernetes-worker
