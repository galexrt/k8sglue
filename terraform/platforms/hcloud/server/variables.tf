variable "hcloud_token" {
  type = "string"
}

variable "ssh_key_name" {
  type = "string"
}

variable "ssh_key_public" {
  type = "string"
}

variable "ssh_key_private" {
  type = "string"
}

variable "dns_domain" {
  default = "example.com"
}

variable "hostname_pattern" {
  type = "string"
}

variable "server_count" {
  type = "string"
}

variable "instance_type" {
  type = "string"
}

variable "cf_email" {}

variable "cf_token" {}
