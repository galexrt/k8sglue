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

variable "hcloud_salt_master_instance_type" {
  type    = "string"
  default = "cx11"
}

variable "hcloud_master_instance_type" {
  type    = "string"
  default = "cx11"
}

variable "hcloud_worker_instance_type" {
  type    = "string"
  default = "cx11"
}

variable "dns_domain" {
  type    = "string"
  default = "example.com"
}

# DNS
variable "cf_email" {
  type = "string"
}

variable "cf_token" {
  type = "string"
}

# Hostnames
variable "hostname_etcd" {
  default = "k8s02-etcd%d"
}

variable "hostname_kubernetes_master" {
  default = "k8s02-master%d"
}

variable "hostname_kubernetes_worker" {
  default = "k8s02-worker%d"
}

variable "hostname_salt_masters" {
  default = "k8s02-salt-master%d"
}

variable "hostname_salt_master" {
  default = "k8s02-salt-master"
}

variable "salt_master_count" {
  type    = "string"
  default = 1
}

variable "master_count" {
  type    = "string"
  default = 3
}

variable "worker_count" {
  type    = "string"
  default = 0
}
