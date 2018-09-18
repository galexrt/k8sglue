{%- if 'act' in data and data['act'] == 'pend' %}
{%- set node = data['id'] %}
{%-   if not salt['file.file_exists']('/etc/salt/pki/master/minions/'+node) %}
call salt-master minion key verify orch:
  runner.state.orchestrate:
    - args:
      - mods: orch.salt-key
      - pillar:
          minion_to_check: '{{ node }}'
{%-   endif %}
{%- endif %}
