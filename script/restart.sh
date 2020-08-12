#!/bin/bash
cd `pwd`
pid=`pidof taskdash`
if [ "$pid" == "" ]; then
    sudo docker-compose -f docker/docker-compose.yml up -d
else
    sudo docker-compose -f docker/docker-compose.yml restart
fi