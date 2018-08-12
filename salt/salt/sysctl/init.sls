{% for key, value in pillar.get('sysctl', {}).items() %}
'sysctl set {{ key }}':
  sysctl.present:
    - name: "{{ key }}"
    - value: "{{ value }}"
    - config: "/etc/sysctl.d/10-custom.conf"
{% endfor %}
