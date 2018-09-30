{% macro get_ip() -%}
  {%- if salt['pillar.get']('cluster_config:network:preferred_ipversion', 4) == 4 -%}
    {{ salt['mine.get']('roles:salt_master', 'ipv4_addresses', tgt_type='grain') }}
  {%- else -%}
    {{ salt['mine.get']('roles:salt_master', 'ipv6_addresses', tgt_type='grain') }}
  {%- endif -%}
{%- endmacro %}
