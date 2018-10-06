# k8sglue salt
Current saltstack files don't really do anything.. This will change in the future.

## Reactor

The Salt code (will) make(s) heavy use of Salt reactor and the experimental [Thorium Complex Reactor](https://docs.saltstack.com/en/latest/topics/thorium/index.html) to keep the infrastructure "reactive" and "dynamic".

### Events

List of events with description and what data they contain:

| Name                            | Description                                                                   | Data                         |
| ------------------------------- | ----------------------------------------------------------------------------- | ---------------------------- |
| `custom/node/need-reboot`       | Indicates that a node needs reboot (e.g., after updates have been installed). | -                            |
| `custom/node/updates-available` | Indicates that a node has updates available.                                  | `count` of updates available |
