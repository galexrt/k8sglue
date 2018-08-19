variable "ssh_key_name" {
  type    = "string"
  default = "Terraform and Salt Deploy"
}

variable "ssh_key_public" {
  type    = "string"
  default = "~/.ssh/id_rsa.pub"
}

variable "ssh_key_private" {
  type    = "string"
  default = "~/.ssh/id_rsa"
}

# Hetzner Cloud ================================================================
variable "hcloud_token" {
  type = "string"
}

variable "hcloud_master_instance_type" {
  type    = "string"
  default = "cx11"
}

variable "hcloud_worker_instance_type" {
  type    = "string"
  default = "cx11"
}

# DNS
variable "dns_domain" {
  type    = "string"
  default = "example.com"
}

variable "cf_email" {
  type = "string"
}

variable "cf_token" {
  type = "string"
}

# Hostnames
variable "hostname_kubernetes_masters" {
  default = "k8s02-masters"
}

variable "hostname_kubernetes_master" {
  default = "k8s02-master%d"
}

variable "hostname_kubernetes_worker" {
  default = "k8s02-worker%d"
}

variable "salt_master_count" {
  type    = "string"
  default = 3
}

variable "master_count" {
  type    = "string"
  default = 3
}

variable "worker_count" {
  type    = "string"
  default = 1
}
