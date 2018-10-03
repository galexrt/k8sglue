{%- set defaultInterface = salt['grains.get']('defaultInterface', '') %}
mine_functions:
  ipv4_addresses:
    - mine_function: network.ip_addrs
    - {{ defaultInterface }}
  ipv6_addresses:
    - mine_function: network.ip_addrs6
    - {{ defaultInterface }}
