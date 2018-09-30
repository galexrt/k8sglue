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
    - refresh: true
    - require:
      - pkgrepo: docker-ce

docker.service override mount propagation:
  file.managed:
    - template: jinja
    - name: /etc/systemd/system/docker.service.d/10-mount-propagation.conf
    - source: salt://crio/templates/etc/systemd/system/docker.service.d/10-mount-propagation.conf
    - user: root
    - group: root
    - mode: '0640'
    - require:
      - pkg: docker-ce
  module.run:
    - name: service.systemctl_reload
    - onchanges:
      - file: 'docker.service override mount propagation'

docker service running:
  service.running:
    - name: docker
    - require:
      - pkg: docker-ce
      - file: /etc/systemd/system/docker.service.d/10-mount-propagation.conf
    - enable: true
