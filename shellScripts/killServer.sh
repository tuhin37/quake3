#! /bin/sh
# PID=$(ps -aux |grep "/bin/bash /bin/sh /app/shellScripts/RunServer.sh" |grep -v grep |sed -e's/  */ /g' |cut -d ' ' -f 2)

#  find PID of 
SCRIPT_PID=$(ps -aux |grep "/bin/sh /app/shellScripts/RunServer.sh" |grep -v grep |sed -e's/  */ /g' |cut -d " " -f 2)
# echo $PID RunServer.sh
if [ "$SCRIPT_PID" = "" ]; then
    echo "skipped" 
else
    kill -9 $SCRIPT_PID
    echo "stopped"
fi

# after killing the shell script, the q3a server continues to run as a zombie process. Lets kill that now
SERVER_PID=$(ps -aux|grep -v "grep" |grep "/app/quake3 ioq3ded.x86_64 +exec server.cfg +exec levels.cfg +exec bots.cfg" |sed -e's/  */ /g' |cut -d " " -f 2)
if [ "$SERVER_PID" != "" ]; then
    kill -9 $SERVER_PID
fi