docker-ce repo:
  pkgrepo.managed:
    - name: docker-ce
    - humanname: docker-ce
    - baseurl: https://download.docker.com/linux/fedora/$releasever/$basearch/stable
    - gpgcheck: 1
    - gpgkey: https://download.docker.com/linux/fedora/gpg

docker-ce package:
  pkg.latest:
    - name: docker-ce
    - refresh: True
    - require:
      - pkgrepo: docker-ce

docker service running:
  service.running:
    - require:
      - pkg: docker-ce
    - enable: True
