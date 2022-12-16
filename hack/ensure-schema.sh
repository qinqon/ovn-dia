#!/bin/bash -xe

OVN_VERSION=v22.06.0
NBDB_SCHEMA=pkg/nbdb/ovn-nb.ovsschema

if [ ! -f $NBDB_SCHEMA ]; then
    curl -sSL https://raw.githubusercontent.com/ovn-org/ovn/${OVN_VERSION}/ovn-nb.ovsschema -o $NBDB_SCHEMA
fi
