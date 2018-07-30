#!/usr/bin/env bash

{% set ttl = salt['pillar.get']('kubernetes').get('kubeadm').get('token').get('ttl', '10m') %}

kubeadm token create --ttl {{ ttl }}
