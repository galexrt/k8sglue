resource "cloudflare_record" "dns_v4" {
  count   = "${length(var.names)}"
  domain  = "${var.domain}"
  value   = "${element(var.addresses_ipv4, count.index)}"
  name    = "${element(var.names, count.index)}"
  type    = "A"
  proxied = "${var.proxied}"
}

resource "cloudflare_record" "dns_v6" {
  count   = "${length(var.names)}"
  domain  = "${var.domain}"
  value   = "${element(var.addresses_ipv6, count.index)}"
  name    = "${element(var.names, count.index)}"
  type    = "AAAA"
  proxied = "${var.proxied}"
}
