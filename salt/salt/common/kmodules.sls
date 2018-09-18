{%- for kmod, opts in salt['pillar.get']('kmods', {}).items() %}
'add and load {{ kmod }}':
{%- if opts.load|default(true) %}
  kmod.present:
{%- else %}
  kmod.absent:
{%- endif %}
    - mods:
      - '{{ kmod }}'
    - persist: False
{%- if not opts.load|default(true) %}
    - comment: False
{%- endif %}
{%- endfor %}

create kmodules list file:
  file.managed:
    - name: /etc/modules-load.d/10-custom.conf
    - user: root
    - group: root
    - mode: '0644'
    - source: salt://common/templates/etc/modules-load.d/10-custom.conf
    - template: jinja
