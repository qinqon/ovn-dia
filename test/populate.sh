#!/bin/bash -xe

ovn-nbctl ls-add hv1
ovn-nbctl set Logical_Switch hv1 other_config:subnet=192.168.50.0/24
ovn-nbctl lsp-add hv1 vm1
ovn-nbctl lsp-set-addresses vm1 "40:44:00:00:00:01 192.168.50.20"
ovn-nbctl lsp-add hv1 hv1-to-join
ovn-nbctl lsp-set-options hv1-to-join "router-port=vms"

ovn-nbctl ls-add hv2
ovn-nbctl set Logical_Switch hv2 other_config:subnet=192.168.51.0/24
ovn-nbctl lsp-add hv2 vm2
ovn-nbctl lsp-set-addresses vm2 "40:44:00:00:00:02 192.168.51.20"
ovn-nbctl lsp-add hv2 hv2-to-join
ovn-nbctl lsp-set-options hv2-to-join "router-port=vms"

ovn-nbctl lr-add join
ovn-nbctl lrp-add join vms "40:44:00:00:00:00" "192.168.50.1/24"
ovn-nbctl show
