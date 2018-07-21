swapoff -a:
  cmd.run

remove swap fstab entries:
  file.comment:
    - name: /etc/fstab
    - regex: ^[^#]*swap.*
