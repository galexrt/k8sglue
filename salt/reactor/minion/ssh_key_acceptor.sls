highstate run:
  local.state.apply:
    - tgt: {{ data['id'] }}
    - require:
      - state: 'safe minion key accept'

kubernetes join token:
  runner.state.orchestrate:
    - args:
      - mods: orch.kubernetes
      - pillar:
          event_tag: {{ tag }}
          event_data: {{ data | json }}
      - require:
        - state: 'highstate run'
