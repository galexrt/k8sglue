{%- from 'glue/utils/ips.sls' import get_ip with context %}
{%- set ip_version = salt['pillar.get']('cluster_config:network:preferred_ipversion', 4) %}
{%- set kubernetes_master_ips = salt['mine.get']('roles:salt_master', 'ipv'+ip_version+'_addresses', tgt_type='grain').values() %}

run k8s master lb container:
  docker_container.running:
    - name: k8s-master-lb
    - image: xetys/k8s-master-lb:latest
    - port_bindings:
      - 127.0.0.1:16443:16443
    - restart_policy: always
    - detach: true
    - cmd: '{{ kubernetes_master_ips }}'
    - shutdown_timeout: 15
