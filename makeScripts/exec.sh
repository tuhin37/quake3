#! /bin/bash
CONTAINER_ID=$(docker ps -a |grep drag/q3asrv |cut -d " " -f 1)
docker exec -it $CONTAINER_ID /bin/sh