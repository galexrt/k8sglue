#!/usr/bin/env bash

{% set host = salt['grains.get']('host', 'default') %}
{% set roles = salt['grains.get']('roles', []) %}
{% set containerRuntime = pillar.get('containerRuntime', "crio") %}
{% set kubernetes_master_ca_cert_hash = salt['mine.get']('roles:kubernetes-master-init', 'kubernetes-master-ca-cert-hash', tgt_type='grain').values()[0] %}
{% set kubernetes_join_token = salt['mine.get']('roles:kubernetes-master-init', 'kubernetes-join-token-'+host) %}

kubeadm join \
{%- if "kubernetes-master" in roles %}
    --master \
{%- endif %}
    --feature-gates=DynamicKubeletConfig \
{%- if containerRuntime == "crio" %}
    --cri-socket=/var/run/crio/crio.sock \
{%- endif %}
    --discovery-token "{{ kubernetes_join_token }}" \
    --discovery-token-ca-cert-hash "sha256:{{ kubernetes_master_ca_cert_hash }}" | \
    tee /var/log/kubeadm-join.log
