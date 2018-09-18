selinux:
  state: permissive
  type: targeted
containerRuntime: crio
kmods:
  br_netfilter:
    load: true
  ip_vs:
    load: true
  ip_vs_rr:
    load: true
  ip_vs_wrr:
    load: true
  ip_vs_sh:
    load: true
  nf_conntrack_ipv4:
    load: true
  nf_conntrack_ipv6:
    load: true
nameservers:
  - 1.1.1.1
  - 8.8.8.8
  - 1.0.0.1
  - 8.4.4.8
sysctl:
  fs.file-max: 2097152
  fs.suid_dumpable: 0
  kernel.dmesg_restrict: 1
  kernel.panic: 10
  kernel.pid_max: 4194303
  kernel.sched_autogroup_enabled: 0
  net.core.somaxconn: 4096
  net.nf_conntrack_max: 1024000
  net.netfilter.nf_conntrack_expect_max: 1024
  net.netfilter.nf_conntrack_max: 1024000
  net.ipv4.conf.all.igmpv2_unsolicited_report_interval: 10000
  net.ipv4.conf.all.igmpv3_unsolicited_report_interval: 1000
  net.ipv4.conf.all.ignore_routes_with_linkdown: 0
  net.ipv4.fwmark_reflect: 0
  net.ipv4.icmp_msgs_burst: 50
  net.ipv4.icmp_msgs_per_sec: 1000
  net.ipv4.ip_forward: 1
  net.ipv4.conf.all.forwarding: 1
  net.ipv4.conf.default.forwarding: 1
  net.ipv6.conf.all.forwarding: 1
  net.ipv6.conf.default.forwarding: 1
  net.ipv4.ip_local_port_range: "8192 65535"
  net.ipv4.ipfrag_secret_interval: 600
  net.ipv4.tcp_tw_reuse: 1
  net.ipv4.tcp_syncookies: 1
  net.ipv4.tcp_syn_retries: 2
  net.ipv4.tcp_synack_retries: 2
  net.ipv4.tcp_max_syn_backlog: 4096
  net.ipv4.conf.all.send_redirects: 0
  net.ipv4.conf.default.send_redirects: 0
  net.ipv4.conf.lo.accept_source_route: 1
  net.ipv4.conf.all.accept_source_route: 0
  net.ipv4.conf.default.accept_source_route: 0
  net.ipv6.conf.all.accept_source_route: 0
  net.ipv6.conf.default.accept_source_route: 0
  net.ipv4.conf.all.rp_filter: 1
  net.ipv4.conf.default.rp_filter: 1
  net.ipv4.conf.all.accept_redirects: 0
  net.ipv4.conf.default.accept_redirects: 0
  net.ipv4.conf.all.secure_redirects: 1
  net.ipv4.conf.default.secure_redirects: 1
  net.ipv6.conf.all.accept_redirects: 0
  net.ipv6.conf.default.accept_redirects: 0
  net.ipv4.conf.all.log_martians: 1
  net.ipv4.conf.default.log_martians: 1
  net.ipv4.tcp_fin_timeout: 7
  net.ipv4.tcp_keepalive_time: 420
  net.ipv4.tcp_keepalive_probes: 5
  net.ipv4.tcp_keepalive_intvl: 25
  net.ipv4.conf.all.bootp_relay: 0
  net.ipv4.conf.all.proxy_arp: 0
  net.ipv4.tcp_timestamps: 1
  net.ipv4.icmp_echo_ignore_all: 0
  net.ipv4.icmp_echo_ignore_broadcasts: 1
  net.ipv4.icmp_ignore_bogus_error_responses: 1
  net.ipv4.tcp_rfc1337: 1
  net.ipv6.ip6frag_secret_interval: 600
  net.ipv6.route.max_size: 16384
  net.ipv6.xfrm6_gc_thresh: 32768
  vm.overcommit_memory: 1
  vm.overcommit_ratio: 50
  vm.panic_on_oom: 0
