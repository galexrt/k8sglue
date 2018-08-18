saltutil sync all:
  local.saltutil.sync_all:
    - tgt: {{ data['id'] }}
