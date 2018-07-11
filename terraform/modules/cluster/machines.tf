resource "null_resource" "kubernetes" {
  count = "${var.server_count}"

  # Changes to any master instance of the cluster requires re-provisioning of masters
  triggers {
    cluster_instance_ids = "${join(",", var.ids)}"
    salt_sha256          = "${data.archive_file.salt.output_base64sha256}"
  }

  # Bootstrap script can run on any instance of the cluster
  # So we just choose the first in this case
  connection {
    host        = "${element(var.addresses_ipv4, count.index)}"
    type        = "ssh"
    user        = "root"
    private_key = "${file(var.ssh_key_private)}"
  }

  provisioner "remote-exec" {
    inline = ["sudo salt-call --local test.ping"]
  }
}

data "template_file" "nodes_salt" {
  count = "${var.server_count}"

  template = <<EOF
- id: '$${id}'
  hostname: '$${hostname}'
  address_ipv4: '$${address_ipv4}'
  address_ipv6: '$${address_ipv6}'
EOF

  vars {
    id           = "${element(var.ids, count.index)}"
    hostname     = "${element(var.names, count.index)}"
    address_ipv4 = "${element(var.addresses_ipv4, count.index)}"
    address_ipv6 = "${element(var.addresses_ipv6, count.index)}"
  }
}

resource "local_file" "nodes_salt" {
  content = <<EOF
nodes:
${join("", data.template_file.nodes_salt.*.rendered)}
EOF

  filename = "${path.root}/salt/nodes-salt.yml"
}
