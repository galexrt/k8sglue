create kubeadm token:
  module.run:
    - name: mine.send
    - func: cmd.run
    - m_name: kubernetes-join-token-{{ salt['pillar.get']('host') }}
    - kwargs:
        cmd: kubeadm token create --ttl=10m
