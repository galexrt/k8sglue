{%- if salt['file.file_exists']('/srv/pillar/config.yaml') or salt['file.file_exists']('/tmp/k8sglue/pillar/config.yaml') %}
{%-     import_yaml "" as cluster_config %}
{%- else %}
cluster_config:
  containerRuntime: crio
  network:
    preferred_ipversion: 6
{%- endif  %}
