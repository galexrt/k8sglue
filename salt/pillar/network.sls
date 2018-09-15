{%- set defaultInterface = salt['grains.get']('defaultInterface', 'eth0') %}
mine_functions:
  ip_address:
    - mine_function: network.ip_addrs
    - {{ defaultInterface }}
