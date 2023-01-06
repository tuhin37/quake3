#! /bin/bash
docker stop $(docker ps -a |grep fidays/quake3 |cut -d " " -f 1)
exit 0