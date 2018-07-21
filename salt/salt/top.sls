{% set containerRuntime = pillar.get('containerRuntime', "crio") -%}
base:
  '*':
    - common
    - sysctl
    - system-maintenance
    - salt-minion
  'G@roles:salt-master':
    - salt-master
  'G@roles:kubernetes-*':
    - kubernetes-basics
{%- if containerRuntime == "crio" %}
    - crio
{%- else %}
    - docker
{%- endif %}
  'G@roles:kubernetes-master':
  'G@roles:kubernetes-worker':
