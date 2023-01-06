#! /bin/bash
CONTAINER_ID=$(docker ps -a |grep fidays/quake3 |cut -d " " -f 1)
if [ "$CONTAINER_ID" = "" ]; then
    # run a new with /bash/sh
    docker run --rm --name quake3 -p 5000:5000/tcp -p 27960:27960/udp --env RAM=128 --env PORT=27960 --env PASSWORD=password --env TOKEN=token -it fidays/quake3
else
    docker exec -it $CONTAINER_ID /bin/sh
fi
