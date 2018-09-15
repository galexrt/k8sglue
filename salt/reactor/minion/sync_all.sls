saltutil sync all:
  local.saltutil.sync_all:
    - tgt: '{{ data['id'] }}'
    - reload_modules: True

saltutil refresh grains:
  local.saltutil.refresh_grains:
    - tgt: '{{ data['id'] }}'

saltutil update mine:
  local.saltutil.runner:
    - name: mine.update
    - arg:
      - '*'
    - tgt: '{{ data['id'] }}'
