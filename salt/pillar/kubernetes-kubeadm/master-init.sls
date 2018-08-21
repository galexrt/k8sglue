mine_functions:
  kubernetes_master_ca_cert_hash:
    - mine_function: cmd.run_stdout
    - {{ salt['grains.get']('shell', '/bin/sh') }} -c "openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'"
    - /root
    - python_shell: true
