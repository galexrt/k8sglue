{% macro get_ips(tgt, tgt_type, out_type) -%}
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
    {%- if out_type == "" or out_type == "yaml" %}
        {{- ips | yaml }}
    {%- elif out_type == "string:comma" %}
        {{- ips | join(',') }}
    {%- elif out_type == "string:space" %}
        {{- ips | join(' ') }}
    {%- endif %}
{%- endmacro %}
