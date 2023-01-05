#! /bin/bash
docker stop $(docker ps -a |grep drag/q3asrv |cut -d " " -f 1)
exit 0