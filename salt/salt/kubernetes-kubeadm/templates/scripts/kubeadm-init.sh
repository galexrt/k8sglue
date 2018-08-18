#!/usr/bin/env bash

{% set defaultInterface = salt['grains.get']('defaultInterface', 'eth0') %}
{% set ipAddress = salt['grains.get']('ip_interfaces')[defaultInterface][0] %}
{% set containerRuntime = salt['pillar.get']('containerRuntime', "crio") %}
{% set host = salt['grains.get']('host') %}

kubeadm init \
    --apiserver-advertise-address={{ ipAddress }} \
{%- if containerRuntime == "crio" %}
    --cri-socket=/var/run/crio/crio.sock \
{%- endif %}
    --node-name={{ host }} \
    --pod-network-cidr=100.64.0.0/13 \
    --service-cidr=100.72.0.0/16 | \
    tee /var/log/kubeadm-init.log
