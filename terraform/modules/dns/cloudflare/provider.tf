provider "cloudflare" {
  email = "${var.cf_email}"
  token = "${var.cf_token}"
}
