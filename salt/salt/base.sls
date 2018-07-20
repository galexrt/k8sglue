{% import_yaml "/srv/salt/config.yml" as config %}
base:
  '*':
    - common
    - salt-minion
