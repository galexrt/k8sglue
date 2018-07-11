resource "hcloud_ssh_key" "default" {
  name       = "${var.ssh_key_name}"
  public_key = "${file(var.ssh_key_public)}"
}
