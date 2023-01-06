#!/bin/bash

IMAGE_ID=$(docker images |grep fidays/quake3 |sed "s/  */ /g" |cut -d ' ' -f 3)
docker rmi $IMAGE_ID
