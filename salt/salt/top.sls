{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
base:
  '*':
    - common
    - salt-minion
    - sysctl
  'G@roles:kubernetes-*':
    - kubernetes-basics
{%- if containerRuntime == "crio" %}
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
