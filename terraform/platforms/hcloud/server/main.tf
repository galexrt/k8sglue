resource "hcloud_server" "servers" {
  count       = "${var.server_count}"
  name        = "${format(var.hostname_pattern, count.index+1)}.${var.dns_domain}"
  server_type = "${var.instance_type}"                                             # use smallest instance with local storage for masters
  image       = "fedora-28"
  datacenter  = "fsn1-dc8"
  ssh_keys    = ["${var.ssh_key_name}"]
  keep_disk   = false
  user_data   = "${var.cloud_config}"

  connection {
    type        = "ssh"
    user        = "root"
    private_key = "${file(var.ssh_key_private)}"
  }

  provisioner "remote-exec" {
    inline = [
      "${var.script}",
    ]
  }
}

module "servers_dns" {
  source = "./../../../modules/dns/cloudflare"

  record_count   = "${var.server_count}"
  cf_email       = "${var.cf_email}"
  cf_token       = "${var.cf_token}"
  domain         = "${var.dns_domain}"
  names          = "${hcloud_server.servers.*.name}"
  addresses_ipv4 = "${hcloud_server.servers.*.ipv4_address}"
  addresses_ipv6 = "${hcloud_server.servers.*.ipv6_address}"
  proxied        = "false"
}
