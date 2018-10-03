{% macro get_ips(tgt, tgt_type) -%}
    {%- if salt['pillar.get']('cluster_config:network:preferred_ipversion', 4) == 4 -%}
        {%- set raw_addresses = salt['mine.get'](tgt, 'ipv4_addresses', tgt_type=tgt_type) %}
    {%- else -%}
        {%- set raw_addresses = salt['mine.get'](tgt, 'ipv6_addresses', tgt_type=tgt_type) %}
    {%- endif -%}
    {%- set raw_addresses = salt['mine.get']('*', 'ipv4_addresses') %}
    {%- if raw_addresses is not defined %}
        {{- raise('No IPs found from minions ' + minion) }}
    {%- endif %}
    {%- set ips = [] %}
    {%- for minion, addresses in raw_addresses.items() %}
        {%- if addresses is not defined %}
            {{- raise('No IPs found for minion ' + minion) }}
        {%- endif %}
        {%- for ip in addresses %}
            {%- do ips.append(ip) %}
        {%- endfor %}
    {%- endfor %}
    {{ ips }}
{%- endmacro %}
