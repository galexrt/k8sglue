output "names" {
  value = "${join(",", hcloud_server.servers.*.name)}"
}

output "ids" {
  value = "${join(",", hcloud_server.servers.*.id)}"
}

output "addresses_ipv4" {
  value = "${join(",", hcloud_server.servers.*.ipv4_address)}"
}

output "addresses_ipv6" {
  value = "${join(",", hcloud_server.servers.*.ipv6_address)}"
}
