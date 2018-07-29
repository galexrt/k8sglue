mine_functions:
  kubernetes-master-ca-cert-hash:
    - mine_function: cmd.run
    - {{ salt['grains.get']('shell', '/bin/sh') }} -c "openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'"
