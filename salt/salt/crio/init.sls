install crio package:
  pkg.latest:
    - name: cri-o
    - refresh: True

/etc/crictl.yaml:
  file.managed:
    - template: jinja
    - source: salt://crio/etc/crictl.yaml
    - user: root
    - group: root
    - mode: '0640'
    - require:
      - pkg: cri-o

/etc/crio/crio.conf:
  file.managed:
    - template: jinja
    - source: salt://crio/etc/crio/crio.conf
    - user: root
    - group: root
    - mode: '0640'
    - require:
      - pkg: cri-o

enabled and start crio service:
  service.running:
    - name: crio
    - enable: True
    - require:
      - file: /etc/crio/crio.conf
      - file: /etc/crictl.yaml
      - pkg: cri-o
