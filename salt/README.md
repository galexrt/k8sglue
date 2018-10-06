# k8sglue salt
Current saltstack files don't really do anything.. This will change in the future.

## Reactor

The Salt code (will) make(s) heavy use of Salt reactor and the experimental [Thorium Complex Reactor](https://docs.saltstack.com/en/latest/topics/thorium/index.html) to keep the infrastructure "reactive" and "dynamic".

### Events

List of events with description and what data they contain:

| Name                               | Description                                                                   | Data                                   |
| ---------------------------------- | ----------------------------------------------------------------------------- | -------------------------------------- |
| `custom/node/need-reboot`          | Indicates that a node needs reboot (e.g., after updates have been installed). | `needed` (bool)                        |
| `custom/node/os/updates-available` | Indicates that a node has package updates available.                          | `count` (int) of updates available     |
| `custom/node/os/updates-done`      | Indicated that a node has installed it's package updates.                     | `status` (int, `0` success, `1` fail) |
