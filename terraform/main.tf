module "hcloud_ssh_key" {
  source = "./platforms/hcloud/ssh_key"

  hcloud_token = "${var.hcloud_token}"

  ssh_key_name   = "${var.ssh_key_name}"
  ssh_key_public = "${var.ssh_key_public}"
}

module "hcloud_kubernetes_masters" {
  source = "./platforms/hcloud/server"

  hcloud_token     = "${var.hcloud_token}"
  cf_email         = "${var.cf_email}"
  cf_token         = "${var.cf_token}"
  ssh_key_name     = "${module.hcloud_ssh_key.name}"
  ssh_key_public   = "${var.ssh_key_public}"
  ssh_key_private  = "${var.ssh_key_private}"
  dns_domain       = "${var.dns_domain}"
  hostname_pattern = "${var.hostname_kubernetes_master}"
  server_count     = "${var.master_count}"
  instance_type    = "${var.hcloud_master_instance_type}"
}

module "hcloud_kubernetes_workers" {
  source = "./platforms/hcloud/server"

  hcloud_token     = "${var.hcloud_token}"
  cf_email         = "${var.cf_email}"
  cf_token         = "${var.cf_token}"
  ssh_key_name     = "${module.hcloud_ssh_key.name}"
  ssh_key_public   = "${var.ssh_key_public}"
  ssh_key_private  = "${var.ssh_key_private}"
  dns_domain       = "${var.dns_domain}"
  hostname_pattern = "${var.hostname_kubernetes_worker}"
  server_count     = "${var.worker_count}"
  instance_type    = "${var.hcloud_worker_instance_type}"
}

module "kubernetes_masters_dns" {
  source = "./modules/dns/cloudflare"

  cf_email       = "${var.cf_email}"
  cf_token       = "${var.cf_token}"
  domain         = "${var.dns_domain}"
  names          = "${list(var.hostname_kubernetes_masters)}"
  addresses_ipv4 = "${split(",", module.hcloud_kubernetes_masters.addresses_ipv4)}"
  addresses_ipv6 = "${split(",", module.hcloud_kubernetes_masters.addresses_ipv6)}"
  proxied        = "false"
}

module "kubernetes_cluster" {
  source = "./modules/cluster"

  ssh_key_private = "${var.ssh_key_private}"
  server_count    = "${var.master_count + var.worker_count}"
  salt_master     = "${var.hostname_kubernetes_masters}"
  ids             = "${concat(compact(split(",", module.hcloud_kubernetes_masters.ids)), compact(split(",", module.hcloud_kubernetes_workers.ids)))}"
  names           = "${concat(compact(split(",", module.hcloud_kubernetes_masters.names)), compact(split(",", module.hcloud_kubernetes_workers.names)))}"
  addresses_ipv4  = "${concat(compact(split(",", module.hcloud_kubernetes_masters.addresses_ipv4)), compact(split(",", module.hcloud_kubernetes_workers.addresses_ipv4)))}"
  addresses_ipv6  = "${concat(compact(split(",", module.hcloud_kubernetes_masters.addresses_ipv6)), compact(split(",", module.hcloud_kubernetes_workers.addresses_ipv6)))}"
}
