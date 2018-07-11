crio package and service:
  pkg.latest:
    - name: cri-o
    - refresh: True
  service.running:
    - require:
      - pkg: cri-o
    - name: crio
    - enable: True
