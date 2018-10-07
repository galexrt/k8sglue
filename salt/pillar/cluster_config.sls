{%- if salt['file.file_exists']('/srv/pillar/cluster_config.yaml') or salt['file.file_exists']('/tmp/k8sglue/pillar/cluster_config.yaml') %}
{%-     import_yaml "cluster_config.yaml" as clusterConfig %}
{%- else %}
clusterConfig:
  containerRuntime: crio
  network:
    preferredIPVersion: 6
{%- endif  %}
