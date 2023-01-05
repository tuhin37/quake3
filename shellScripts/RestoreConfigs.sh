#! /bin/sh

# delete existing cfg files
rm /app/quake3/baseq3/autoexec.cfg
rm /app/quake3/baseq3/bots.cfg
rm /app/quake3/baseq3/levels.cfg
rm /app/quake3/baseq3/server.cfg

# restore new cfg files from backup
cp /app/quake3/baseq3/autoexec.bak /app/quake3/baseq3/autoexec.cfg
cp /app/quake3/baseq3/bots.bak /app/quake3/baseq3/bots.cfg
cp /app/quake3/baseq3/levels.bak /app/quake3/baseq3/levels.cfg
cp /app/quake3/baseq3/server.bak /app/quake3/baseq3/server.cfg


echo "SH | configs restored"