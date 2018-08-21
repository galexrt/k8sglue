{% set defaultInterface = salt['grains.get']('defaultInterface') %}
mine_functions:
  ip_address:
    - mine_function: network.ip_addrs
    - {{ defaultInterface }}
