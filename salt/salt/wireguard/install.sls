enable copr wireguard repo:
  pkgrepo.managed:
    - name: jdoss-wireguard
    - humanname: Copr repo for wireguard owned by jdoss
    - baseurl: https://copr-be.cloud.fedoraproject.org/results/jdoss/wireguard/fedora-$releasever-$basearch/
    - enabled: 1
    - gpgcheck: 1
    - gpgkey: https://copr-be.cloud.fedoraproject.org/results/jdoss/wireguard/pubkey.gpg
#   - repo_gpgcheck: 0
#   - enabled_metadata: 1

install wireguard dkms and tools:
  pkg.latest:
    - pkgs:
      - wireguard-dkms
      - wireguard-tools
    - require:
      - pkgrepo: 'enable copr wireguard repo'
