{%- from 'glue/utils/ips.sls' import get_ips with context %}
{%- set kubernetes_master_ips = get_ips('roles:salt_master', 'grain') %}

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
