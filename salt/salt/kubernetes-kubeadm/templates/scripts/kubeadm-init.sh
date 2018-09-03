#!/usr/bin/env bash

{% set defaultInterface = salt['grains.get']('defaultInterface', 'eth0') %}
{% set ipAddress = salt['grains.get']('ip_interfaces')[defaultInterface]|first %}
{% set containerRuntime = salt['pillar.get']('containerRuntime', "crio") %}
{% set host = salt['grains.get']('host') %}

set -e
set -o pipefail

kubeadm init \
    --apiserver-advertise-address={{ ipAddress }} \
{%- if containerRuntime == "crio" %}
    --cri-socket=/var/run/crio/crio.sock \
{%- endif %}
    --node-name={{ host }} \
    --pod-network-cidr=100.64.0.0/13 \
    --service-cidr=100.72.0.0/16 | \
    tee /var/log/kubeadm-init.log


TOKEN=$(kubeadm token list | tail -n +2 | head -n -1 | awk '{print $1}') || echo "Error getting kubeadm token list from master server. Out: '$TOKEN', Ret: $?"
if [ -n "${TOKEN}" ]; then
    kubeadm token delete "${TOKEN}"
else
    echo "No kubeadm token to delete."
fi

exit 0
