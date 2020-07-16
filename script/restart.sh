#!/bin/bash
cd `pwd`
pid=`pidof ./taskdash`
cp -f nohup.out nohup.out.1
cat /dev/null > nohup.out
if [ "$pid" == "" ]; then
    nohup ./taskdash &
else
    kill -SIGUSR2 ${pid}
fi