{%- if salt['file.file_exists']('/srv/pillar/salt_master_addresses.yaml') or salt['file.file_exists']('/tmp/k8sglue/pillar/salt_master_addresses.yaml') %}
{%-     import_yaml "salt_master_addresses.yaml" as salt_master_addresses %}
{{ salt_master_addresses|yaml(false) }}
{%- else %}
salt:
  master:
    addresses:
      - 127.0.0.1
{%- endif %}
