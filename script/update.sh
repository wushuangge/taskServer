#!/bin/bash
git pull && rm -f videoServer && make && ./restart.sh
