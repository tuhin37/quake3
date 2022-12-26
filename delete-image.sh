#!/bin/bash

IMAGE_ID = `docker images |grep drag/q3asrv |sed "s/  */ /g" |cut -d ' ' -f 3)`
echo $IMAGE_ID
#docker rmi $IMAGE_ID
