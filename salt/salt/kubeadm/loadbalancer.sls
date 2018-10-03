{%- from 'glue/macros/get_ips.sls' import get_ips with context %}
{%- set kubernetes_master_ips = get_ips('roles:kubernetes_master', 'grain', 'string:space') %}

include:
- docker

install python docker package:
  pkg.latest:
    - pkgs:
      - python2-docker
      - python3-docker
    - reload_modules: true

run k8s master lb container:
  docker_container.running:
    - name: k8s-master-lb
    - image: xetys/k8s-master-lb:latest
    - port_bindings:
      - 127.0.0.1:16443:16443
    - restart_policy: always
    - detach: true
    - cmd: '{{ kubernetes_master_ips | yaml_dquote }}'
    - shutdown_timeout: 15
    - require:
      - pkg: 'install python docker package'
      - service: 'start docker service'
