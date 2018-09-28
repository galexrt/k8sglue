#!/usr/bin/env bash

{%- set ttl = salt['pillar.get']('kubernetes:kubeadm:token:ttl', '10m') %}

set -e
set -o pipefail

kubeadm token create --kubeconfig /etc/kubernetes/admin.conf --ttl {{ ttl }}