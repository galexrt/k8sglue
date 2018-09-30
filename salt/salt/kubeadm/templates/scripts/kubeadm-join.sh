#!/usr/bin/env bash

{%- set defaultInterface = salt['grains.get']('defaultInterface', 'eth0') %}
{%- set ipAddress = salt['grains.get']('ip_interfaces')[defaultInterface]|first %}
{%- set containerRuntime = salt['pillar.get']('cluster_config:containerRuntime', 'crio') %}
{%- set roles = salt['grains.get']('roles') %}
{%- set host = salt['grains.get']('host') %}

set -e
set -o pipefail

if [ -z "${KUBEADM_JOIN_TOKEN}" ]; then
    echo "No KUBEADM_JOIN_TOKEN env var given."
    exit 1
fi

KUBERNETES_CA_CERT_HASH="$(openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //')"

{# Move after `kubeadm join` line #}
{%- if "kubernetes_master" in roles %}
#    --experimental-control-plane \
{%- endif %}
kubeadm join \
    --feature-gates=DynamicKubeletConfig=true \
{%- if containerRuntime == "crio" %}
    --cri-socket=/var/run/crio/crio.sock \
{%- endif %}
    --node-name "{{ host }}" \
    --token "${KUBEADM_JOIN_TOKEN}" \
    --discovery-token-ca-cert-hash "sha256:${KUBERNETES_CA_CERT_HASH}" \
    127.0.0.1:16443 | \
    tee /var/log/kubeadm-join.log
