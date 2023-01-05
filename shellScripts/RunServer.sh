#! /bin/sh

# for testing use the long loop
# for VARIABLE in 1 2 3 4 5 6 7 8 9 10
# do
#     echo $VARIABLE
#     sleep 60
# done

/app/quake3/ioq3ded.x86_64 +exec server.cfg +exec levels.cfg +exec bots.cfg
exit 0