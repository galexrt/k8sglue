highstate_run:
  local.state.apply:
    - tgt: {{ data['id'] }}
    - tgt_type: list
