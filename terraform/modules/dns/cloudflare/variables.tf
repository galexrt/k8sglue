variable "cf_email" {}

variable "cf_token" {}

variable "domain" {
  default = "example.com"
}

variable "names" {
  type = "list"
}

variable "addresses_ipv4" {
  type = "list"
}

variable "addresses_ipv6" {
  type = "list"
}

variable "proxied" {
  type    = "string"
  default = false
}
