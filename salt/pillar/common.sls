selinux:
  state: enforcing
  type: targeted
containerRuntime: crio
nameservers:
  - 1.1.1.1
  - 8.8.8.8
  - 1.0.0.1
  - 8.4.4.8