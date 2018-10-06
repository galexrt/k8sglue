'schedule 12-hourly update check':
  schedule.present:
    - enabled: true
    - persist: true
    - maxrunning: 1
    - function: state.sls
    - job_args:
      - system-maintenance.update_check
    - hours: 12
    - splay: 300
    - skip_during_range:
        - start: 10pm
        - end: 6am
