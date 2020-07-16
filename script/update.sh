#!/bin/bash
git pull && rm -f taskdash && make && ./restart.sh
