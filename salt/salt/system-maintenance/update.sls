update os packages:
  pkg.uptodate:
    - refresh: true
    - require_in:
      - event: 'send custom/node/os/updates-done success event'
    - onfail:
      - event: 'send custom/node/os/updates-available event'

send custom/node/os/updates-done success event:
  event.send:
    - name: custom/node/os/updates-available
    - data:
        status: 0

send custom/node/os/updates-done fail event:
  event.send:
    - name: custom/node/os/updates-available
    - data:
        status: 1
