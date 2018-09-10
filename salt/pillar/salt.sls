{#
TODO feed `salt:master:addresses` as a config yaml automatically to the salt-master(s)
TODO the file should automatically be updated when a new node is joined by "just a salt-minion with the salt-master role set"
#}
salt:
  master:
    addresses:
      - 127.0.0.1
