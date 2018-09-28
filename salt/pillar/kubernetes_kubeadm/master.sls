mine_functions:
  kubernetes_master_joined_cluster:
    - mine_function: file.file_exists
    - /etc/kubernetes/manifests/kube-apiserver.yaml

kubernetes:
  kubeadm:
    token:
      ttl: "10m"