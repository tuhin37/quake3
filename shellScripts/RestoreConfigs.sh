#! /bin/sh

# delete existing cfg files
rm $PWD/quake3/baseq3/autoexec.cfg
rm $PWD/quake3/baseq3/bots.cfg
rm $PWD/quake3/baseq3/levels.cfg
rm $PWD/quake3/baseq3/server.cfg

# restore new cfg files from backup
cp $PWD/quake3/baseq3/autoexec.bak $PWD/quake3/baseq3/autoexec.cfg
cp $PWD/quake3/baseq3/bots.bak $PWD/quake3/baseq3/bots.cfg
cp $PWD/quake3/baseq3/levels.bak $PWD/quake3/baseq3/levels.cfg
cp $PWD/quake3/baseq3/server.bak $PWD/quake3/baseq3/server.cfg

echo "SH | configs restored"