mine_functions:
  kubernetes-master-ips:
    - mine_function: network.ip_addrs
    - interface: {{ salt['grains.get']('default_interface', 'eth0') }}
