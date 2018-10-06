{%- set available_pkgs_upgrades = salt['pkg.list_upgrades']()|length %}
{%- if available_pkgs_upgrades > 0 %}
send custom/node/updates-available event:
  event.send:
    - name: custom/node/updates-available
    - data:
        count: {{ available_pkgs_upgrades }}
{%- endif %}
