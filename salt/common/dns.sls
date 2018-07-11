/etc/resolv.conf:
  file.managed:
    - user: root
    - group: root
    - mode: 755
    - source: salt://common/etc/resolv.conf

/etc/dhcp/dhclient-enter-hooks.d:
  file.directory:
    - user: root
    - group: root
    - mode: 755
    - makedirs: True

/etc/dhcp/dhclient-enter-hooks.d/00-nodnsupdate:
  file.managed:
    - name: /etc/dhcp/dhclient-enter-hooks.d/00-nodnsupdate
    - user: root
    - group: root
    - mode: 755
    - template: jinja
    - source: salt://common/etc/dhcp/dhclient-enter-hooks.d/00-nodnsupdate
    - require:
      - file: /etc/dhcp/dhclient-enter-hooks.d
