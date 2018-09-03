{% if 'act' in data and data['act'] == 'pend' %}
call salt-master minion key verify orch:
  runner.state.orchestrate:
    - args:
      - mods: orch.salt-key
      - pillar:
          minion_to_check: {{ data['id'] }}
{% endif %}
