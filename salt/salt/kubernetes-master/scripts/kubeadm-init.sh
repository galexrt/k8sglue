#!/bin/bash
{% set ipaddress = salt['grains.get']('ip_interfaces') -%}

kubeadm init \
    --apiserver-advertise-address={{ ipaddress['eth1'][0] }} \
    --pod-network-cidr=100.64.0.0/13 \
    --service-cidr=100.72.0.0/16 \
    --feature-gates=DynamicKubeletConfig=true \
    {% if pillar.get('containerRuntime', "crio") == "crio" %}--cri-socket=/var/run/crio/crio.sock \{% endif -%}
