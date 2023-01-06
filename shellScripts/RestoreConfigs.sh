#! /bin/sh

AUTOEXEC_BAK=$PWD/config/autoexec.bak
AUTOEXEC=$PWD/ioquake3/baseq3/autoexec.cfg
BOTS_BAK=$PWD/config/bots.bak 
BOTS=$PWD/ioquake3/baseq3/bots.cfg
LEVELS_BACK=$PWD/config/levels.bak
LEVELS=$PWD/ioquake3/baseq3/levels.cfg
SERVER_BAK=$PWD/config/server.bak 
SERVER=$PWD/ioquake3/baseq3/server.cfg


# delete existing cfg files 
if [ -f $SERVER ]; then  
rm $SERVER  
fi  

if [ -f $LEVELS ]; then  
rm $LEVELS  
fi  

if [ -f $BOTS ]; then  
rm $BOTS  
fi  

if [ -f $AUTOEXEC ]; then  
rm $AUTOEXEC  
fi  

# rm $PWD/ioquake3/baseq3/autoexec.cfg
# rm $PWD/ioquake3/baseq3/bots.cfg
# rm $PWD/ioquake3/baseq3/levels.cfg
# rm $PWD/ioquake3/baseq3/server.cfg

# restore new cfg files from backup
cp $AUTOEXEC_BAK $AUTOEXEC
cp $SERVER_BAK $SERVER
cp $LEVELS_BACK $LEVELS
cp $BOTS_BAK $BOTS

echo "SH | configs restored"