#!/bin/bash

target=apisvr

service_PID=$(ps -ef | grep ${target} | grep -v grep | grep -v scp | awk '{print $2}')

if [ -z "${service_PID}" ]; then
    echo "..."
else
    echo "kill ${service_PID}......."
    kill ${service_PID}
    sleep 10
    echo "done"
fi
