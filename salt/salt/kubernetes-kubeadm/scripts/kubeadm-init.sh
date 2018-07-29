#!/usr/bin/env bash
{% set ipaddress = salt['grains.get']('ip_interfaces') -%}

kubeadm init \
    --apiserver-advertise-address={{ ipaddress[salt['grains.get']('default_interface', 'eth0')][0] }} \
{%- if pillar.get('containerRuntime', "crio") == "crio" %}
    --cri-socket=/var/run/crio/crio.sock \
{%- endif %}
    --pod-network-cidr=100.64.0.0/13 \
    --service-cidr=100.72.0.0/16 | \
    tee /var/log/kubeadm-init.log
