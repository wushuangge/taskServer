#!/bin/bash
cd `pwd`

sudo docker-compose -f docker/nsq/docker-compose.yml up -d
