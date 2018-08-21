#!/usr/bin/env bash

{% set defaultInterface = salt['grains.get']('defaultInterface', 'eth0') %}
{% set ipAddress = salt['grains.get']('ip_interfaces')[defaultInterface]|first %}
{% set containerRuntime = salt['pillar.get']('containerRuntime', 'crio') %}
{% set roles = salt['grains.get']('roles') %}
{% set host = salt['grains.get']('host') %}
{% set kubernetes_master_ca_cert_hash = salt['mine.get']('roles:kubernetes-master-init', 'kubernetes_master_ca_cert_hash', tgt_type='grain').values()|first %}
{# TODO Change from master-init to master #}
{% set kubernetes_master_ip = salt['mine.get']('roles:kubernetes-master-init', 'ip_address', tgt_type='grain').values()|random|first %}

set -e
set -o pipefail

if [ -z "${KUBEADM_JOIN_TOKEN}" ]; then
    echo "No KUBEADM_JOIN_TOKEN env var given."
    exit 1
fi

echo "$KUBEADM_JOIN_TOKEN"

kubeadm join \
{%- if "kubernetes-master" in roles %}
    --master \
{%- endif %}
    --feature-gates=DynamicKubeletConfig=true \
{%- if containerRuntime == "crio" %}
    --cri-socket=/var/run/crio/crio.sock \
{%- endif %}
    --node-name "{{ host }}" \
    --token "$KUBEADM_JOIN_TOKEN" \
    --discovery-token-ca-cert-hash "sha256:{{ kubernetes_master_ca_cert_hash }}" \
    {{ kubernetes_master_ip }}:6443 | \
    tee /var/log/kubeadm-join.log
