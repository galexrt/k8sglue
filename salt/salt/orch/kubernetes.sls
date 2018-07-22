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
    - require:
      - salt: setup first kubernetes master

setup kubernetes worker:
  salt.state:
    - tgt: 'roles:kubernetes-worker'
    - tgt_type: grain
    - require:
      - salt: setup kubernetes master
    - sls:
      - kubernetes-worker.kubeadm-join
