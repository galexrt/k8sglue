# TODO Add deb based support (probably needs to add repo for deb)
crio package and service:
  pkg.latest:
    - name: cri-o
    - refresh: True
  service.running:
    - require:
      - pkg: cri-o
    - name: crio
    - enable: True
