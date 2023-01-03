#! /bin/bash
PID=$(ps -aux |grep "/bin/bash ./shellScripts/RunServer.sh" |grep -v grep |sed -e's/  */ /g' |cut -d ' ' -f 2)
# echo $PID
if [ "$PID" = "" ]; then
    echo "skipped" 
else
    kill -9 $PID
    echo "stopped"
fi

