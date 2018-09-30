/etc/resolv.conf:
  file.managed:
    - user: root
    - group: root
    - mode: 755
    - template: jinja
    - source: salt://common/templates/etc/resolv.conf

/etc/dhcp/dhclient-enter-hooks.d:
  file.directory:
    - user: root
    - group: root
    - mode: 755
    - makedirs: true

/etc/dhcp/dhclient-enter-hooks.d/00-nodnsupdate:
  file.managed:
    - user: root
    - group: root
    - mode: 755
    - template: jinja
    - source: salt://common/templates/etc/dhcp/dhclient-enter-hooks.d/00-nodnsupdate
    - require:
      - file: /etc/dhcp/dhclient-enter-hooks.d
