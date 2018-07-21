install crio package:
  pkg.latest:
    - name: cri-o
    - refresh: True

enabled and start crio service:
  service.running:
    - name: crio
    - enable: True
    - require:
      - pkg: cri-o
