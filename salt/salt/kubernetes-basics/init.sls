# TODO Add Kubernetes repo and install kubelet, kubeadm and kubectl (?)
kubernetes repo:
  pkgrepo.managed:
    - name: kubernetes
    - humanname: kubernetes
    - baseurl: https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
    - gpgcheck: 1
    - gpgkey: https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
    - require_in:
      - pkg: kubelet
      - pkg: kubeadm
      - pkg: kubectl

kubelet package:
  pkg.latest:
    - name: kubelet
    - refresh: True
    - require:
      - pkg_repo: kubernetes

kubeadm package:
  pkg.latest:
    - name: kubeadm
    - refresh: True
    - require:
      - pkg_repo: kubernetes

kubectl package:
  pkg.latest:
    - name: kubectl
    - refresh: True
    - require:
      - pkg_repo: kubernetes
