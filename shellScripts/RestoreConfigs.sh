#! /bin/bash

# delete existing cfg files
rm ./config/autoexec.cfg
rm ./config/bots.cfg
rm ./config/levels.cfg
rm ./config/server.cfg

# restore new cfg files from backup
cp ./config/autoexec.bak ./config/autoexec.cfg
cp ./config/bots.bak ./config/bots.cfg
cp ./config/levels.bak ./config/levels.cfg
cp ./config/server.bak ./config/server.cfg


echo "SH | configs restored"