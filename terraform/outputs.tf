output "servers" {
  value = {
      "ids" = "${concat(compact(split(",", module.hcloud_salt_masters.ids)), compact(split(",", module.hcloud_kubernetes_masters.ids)), compact(split(",", module.hcloud_kubernetes_workers.ids)))}"
      "names" = "${concat(compact(split(",", module.hcloud_salt_masters.names)), compact(split(",", module.hcloud_kubernetes_masters.names)), compact(split(",", module.hcloud_kubernetes_workers.names)))}"
      "addresses_ipv4" = "${concat(compact(split(",", module.hcloud_salt_masters.addresses_ipv4)), compact(split(",", module.hcloud_kubernetes_masters.addresses_ipv4)), compact(split(",", module.hcloud_kubernetes_workers.addresses_ipv4)))}"
      "addresses_ipv6" = "${concat(compact(split(",", module.hcloud_salt_masters.addresses_ipv6)), compact(split(",", module.hcloud_kubernetes_masters.addresses_ipv6)), compact(split(",", module.hcloud_kubernetes_workers.addresses_ipv6)))}"
  }
}
