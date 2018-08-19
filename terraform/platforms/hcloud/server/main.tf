resource "hcloud_server" "servers" {
  count       = "${var.server_count}"
  name        = "${format(var.hostname_pattern, count.index+1)}.${var.dns_domain}"
  server_type = "${var.instance_type}"                                             # use smallest instance with local storage for masters
  image       = "fedora-28"
  datacenter  = "fsn1-dc8"
  ssh_keys    = ["${var.ssh_key_name}"]
  keep_disk   = false

  connection {
    type        = "ssh"
    user        = "root"
    private_key = "${file(var.ssh_key_private)}"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo dnf -y install python curl salt-master salt-minion",
      "sudo salt-key --gen-keys=${self.name} --gen-keys-dir=/etc/salt/pki/minion --keysize=4096",
      "sudo mv /etc/salt/pki/minion/${self.name}.pem /etc/salt/pki/minion/minion.pem",
      "sudo mv /etc/salt/pki/minion/${self.name}.pub /etc/salt/pki/minion/minion.pub",
    ]
  }
}

module "servers_dns" {
  source = "./../../../modules/dns/cloudflare"

  cf_email       = "${var.cf_email}"
  cf_token       = "${var.cf_token}"
  domain         = "${var.dns_domain}"
  names          = "${hcloud_server.servers.*.name}"
  addresses_ipv4 = "${hcloud_server.servers.*.ipv4_address}"
  addresses_ipv6 = "${hcloud_server.servers.*.ipv6_address}"
  proxied        = "false"
}
