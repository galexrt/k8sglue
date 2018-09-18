mine_functions:
  kubernetes_worker_joined_cluster:
    - mine_function: file.file_exists
    - /etc/kubernetes/kubelet.conf
