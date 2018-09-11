{%- if 'act' in data and data['act'] == 'pend' %}
{%-   if not salt['file.file_exists']('/etc/salt/pki/master/minions/'+data['id']) %}
call salt-master minion key verify orch:
  runner.state.orchestrate:
    - args:
      - mods: orch.salt-key
      - pillar:
          minion_to_check: {{ data['id'] }}
{%-   endif %}
{%- endif %}
