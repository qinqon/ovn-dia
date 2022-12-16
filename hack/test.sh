#!/bin/bash -xe

tag=ovn-dia-test
nbdb_name=ovn-dia-nbdb

trap "docker stop $nbdb_name" err exit
pushd test

docker build . -t $tag
docker run --rm --name $nbdb_name -d -v /tmp/ovn:/var/run/ovn/ $tag /usr/share/ovn/scripts/ovn-ctl run_nb_ovsdb
docker logs $nbdb_name

sleep 2

docker exec $nbdb_name /populate.sh

sudo ../.out/ovn-dia -nb unix:/tmp/ovn/ovnnb_db.sock



