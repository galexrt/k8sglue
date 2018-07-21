setup first kubernetes master:
  salt.state:
    - tgt: 'roles:kubernetes-master-init'
    - tgt_type: grain
    - sls:
      - kubernetes-master.kubeadm-init

setup kubernetes master:
  salt.state:
    - tgt: 'roles:kubernetes-master'
    - tgt_type: grain
    - sls:
      - kubernetes-master.kubeadm-join

setup kubernetes worker:
  salt.state:
    - tgt: 'roles:kubernetes-worker'
    - tgt_type: grain
    - require:
      - salt: setup_kubernetes_master
    - sls:
      - kubernetes-worker.kubeadm-join
