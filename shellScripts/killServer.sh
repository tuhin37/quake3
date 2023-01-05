#! /bin/sh


# find PID of 
SCRIPT_PID=$(ps -aux |grep -v grep |grep RunServer |sed -e's/  */ /g' |cut -d " " -f 2)
# echo $PID RunServer.sh
if [ "$SCRIPT_PID" = "" ]; then
    echo "skipped" 
else
    kill -9 $SCRIPT_PID
    echo "stopped"
fi



# after killing the shell script, the q3a server continues to run as a zombie process. Lets kill that now
SERVER_PID=$(ps -aux |grep -v grep |grep ioq3ded.x86_64 |sed -e's/  */ /g' |cut -d " " -f 2)
if [ "$SERVER_PID" != "" ]; then
    kill -9 $SERVER_PID
fi