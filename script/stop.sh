#!/bin/bash
cd `pwd`
pid=`pidof taskdash`
if [ "$pid" == "" ]; then
    echo not start
else
    sudo docker-compose -f docker/docker-compose.yml stop
fi