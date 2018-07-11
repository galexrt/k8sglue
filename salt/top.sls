{% import_yaml "config.yml" as config %}
base:
  '*':
    - common
    - salt-minion
  'role:kubernetes':
    - kubernetes-basics
{%- if config.container_runtime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
  'role:salt-master':
    - salt-master
  'role:kubernetes-master':
    - kubernetes-master
  'role:kubernetes-worker':
    - kubernetes-worker
