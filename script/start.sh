#!/bin/bash
cd `pwd`
pid=`pidof ./videoServer`
cp -f nohup.out nohup.out.1
cat /dev/null > nohup.out
if [ "$pid" == "" ]; then
    nohup ./videoServer &
else
    kill -SIGUSR2 ${pid}
fi