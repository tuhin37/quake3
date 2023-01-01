#! /bin/bash
PID=$(ps -aux |grep "/bin/bash ./dummy.sh" |grep -v auto |grep S+ |sed -e's/  */ /g' |cut -d ' ' -f 2)
kill -9 $PID
