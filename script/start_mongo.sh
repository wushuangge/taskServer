#!/bin/bash
cd `pwd`

sudo docker-compose -f docker/mongo/docker-compose.yml up -d
