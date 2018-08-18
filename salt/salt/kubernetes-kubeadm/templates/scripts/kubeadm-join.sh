#!/usr/bin/env bash

{% set defaultInterface = salt['grains.get']('defaultInterface', 'eth0') %}
{% set ipAddress = salt['grains.get']('ip_interfaces')[defaultInterface][0] %}
{% set containerRuntime = salt['pillar.get']('containerRuntime', "crio") %}
{% set roles = salt['grains.get']('roles', []) %}
{% set containerRuntime = pillar.get('containerRuntime', "crio") %}
{% set host = salt['grains.get']('host') %}
{% set master_ca_cert_hash = salt['mine.get']('roles:kubernetes-master-init', 'kubernetes-master-ca-cert-hash', tgt_type='grain').values()[0] %}
{% set join_token = salt['pillar.get']('join-token', None) %}

kubeadm join \
{%- if "kubernetes-master" in roles %}
    --master \
{%- endif %}
    --feature-gates=DynamicKubeletConfig \
{%- if containerRuntime == "crio" %}
    --cri-socket=/var/run/crio/crio.sock \
{%- endif %}
    --node-name={{ host }} \
    --discovery-token "{{ kubernetes_join_token }}" \
    --discovery-token-ca-cert-hash "sha256:{{ kubernetes_master_ca_cert_hash }}" | \
    tee /var/log/kubeadm-join.log
