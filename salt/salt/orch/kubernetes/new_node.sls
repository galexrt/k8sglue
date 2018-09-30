{#
  TODO use salt mine to see if the node has already joined.
  Join scripts must be indepotent!
#}
send custom kubernetes want-to-join event:
  salt.runner:
    - name: event.send
    - tag: custom/kubernetes/want-to-join
    - data:
        minion_to_join: '{{ minion_to_join }}'
