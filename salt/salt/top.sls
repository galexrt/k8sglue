{% import_yaml "/srv/salt/config.yml" as config %}
base:
  '*':
    - common
    - salt-minion
  'G@roles:kubernetes-*':
    - kubernetes-basics
{%- if config.containerRuntime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
  'G@roles:salt-master':
    - salt-master
  'G@roles:kubernetes-master':
    - kubernetes-master
  'G@roles:kubernetes-worker':
    - kubernetes-worker
