{%- if salt['file.file_exists']('/srv/pillar/cluster_config.yaml') or salt['file.file_exists']('/tmp/k8sglue/pillar/cluster_config.yaml') %}
{%-     import_yaml "cluster_config.yaml" as cluster_config %}
{%- else %}
cluster_config:
  containerRuntime: crio
  network:
    preferred_ipversion: 6
{%- endif  %}
