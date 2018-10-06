# https://docs.saltstack.com/en/latest/topics/thorium/index.html
foo:
  reg.list:
    - add: bar
    - match: custom/
    - stamp: True
    - prune: 25
