resource "null_resource" "kubernetes" {
  count = "${var.server_count}"

  # Changes to any master instance of the cluster requires re-provisioning of masters
  triggers {
    cluster_instance_ids = "${join(",", var.ids)}"
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
