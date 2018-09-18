turn swap off:
  cmd.run:
    - name: swapoff -a

remove swap fstab entries:
  file.comment:
    - name: /etc/fstab
    - regex: ^[^#]*swap.*
