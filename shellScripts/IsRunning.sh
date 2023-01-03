#! /bin/bash 
PID=$(ps -aux |grep "/bin/bash ./shellScripts/RunServer.sh" |grep -v grep |sed -e's/  */ /g' |cut -d ' ' -f 2)

if [ "$PID" = "" ]; then
    echo "stopped"
else
    echo "running"
fi

exit 0