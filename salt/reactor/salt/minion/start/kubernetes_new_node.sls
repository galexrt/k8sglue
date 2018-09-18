{%- set minion_to_join = data['id'] %}
call salt-master minion key verify orch:
  runner.state.orchestrate:
    - args:
      - mods: orch.kubernetes.new_node
      - pillar:
          minion_to_join: {{ minion_to_join }}
