kubectl package:
  pkg.latest:
    - name: kubectl
    - refresh: True
    - require:
      - pkgrepo: kubernetes

# TODO Run `kubeadm join --master` for master ((hopefully) available in K8S 1.12)
