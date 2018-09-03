#!/usr/bin/env bash

{%- set defaultInterface = salt['grains.get']('defaultInterface', 'eth0') %}
{%- set ipAddress = salt['grains.get']('ip_interfaces')[defaultInterface]|first %}
{%- set containerRuntime = salt['pillar.get']('containerRuntime', 'crio') %}
{%- set roles = salt['grains.get']('roles') %}
{%- set host = salt['grains.get']('host') %}
{# TODO Change from master-init to master #}
{%- set kubernetes_master_ca_cert_hash = salt['mine.get']('roles:kubernetes-master-init', 'kubernetes_master_ca_cert_hash', tgt_type='grain').values()|first %}
{%- set kubernetes_master_address = salt['mine.get']('roles:kubernetes-master-init', 'ip_address', tgt_type='grain').values()|random|first %}
{%- if kubernetes_master_address is none or kubernetes_master_address == "" %}
{%-  set kubernetes_master_address = salt['pillar.get']('kubernetes:kubeadm:master_address') %}
{%- endif %}
{%- if kubernetes_master_address is none %}
{%-  set kubernetes_master_address = '127.0.0.1:6443' %}
{%- endif %}

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
    {{ kubernetes_master_address }}:6443 | \
    tee /var/log/kubeadm-join.log
