mine_functions:
  salt-master-ips:
    - mine_function: network.ip_addrs
    - interface: {{ salt['grains.get']('defaultInterface', 'eth0') }}
