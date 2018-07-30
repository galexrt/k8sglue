mine_functions:
  ip_address:
    - mine_function: network.ip_addrs
    - interface: {{ salt['grains.get']('defaultInterface', 'eth0') }}
