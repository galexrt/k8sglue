data "archive_file" "salt" {
  type        = "zip"
  source_dir  = "${path.root}/salt"
  output_path = "${path.root}/.tmp/salt.zip"
}
