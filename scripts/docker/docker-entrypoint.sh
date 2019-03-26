#!/bin/sh
set -e

if [ "$1" = 'puti' ]; then
    BASE_DATA=/data/puti
    BASE_APPPUTI=/app/puti
    INIT_PATH=/app/init

    for f in /configs /theme /uploads; do
        if ! test -d ${BASE_DATA}$f; then
            echo "${BASE_DATA}$f is not exist, creating by copy"
            mkdir -p ${BASE_DATA}$f
        fi

        if [ "$(ls -A ${BASE_DATA}$f)" ]; then
            echo "${BASE_DATA}$f is not empty, continue without initialization"
        else
            echo "${BASE_DATA}$f is empty, initialize now"
            cp -r ${INIT_PATH}$f/* ${BASE_DATA}$f/
        fi  

        # Link volumed data with app data
        if ! test -L ${BASE_APPPUTI}$f; then
            echo "link ${BASE_APPPUTI}$f is not exist, creating from source: ${BASE_DATA}$f"
            ln -sfn ${BASE_DATA}$f ${BASE_APPPUTI}$f
        fi
    done

    # logs path
    if ! test -d /data/logs/puti; then
        echo "/data/logs/puti is not exist, creating"
        mkdir -p /data/logs/puti
    fi

    if ! test -L /app/puti/logs; then
        echo "link /app/puti/logs is not exist, creating from source: /data/logs/puti"
        ln -sfn /data/logs/puti /app/puti/logs
    fi

    # Only for the first time
    if ! test -d /app/puti/configs/config.yaml; then
        cp /app/puti/configs/config.yaml.example /app/puti/configs/config.yaml
    fi

    chown -R putiuser:putiuser /data/puti /data/logs/puti
    chmod 0755 -R /data

    exec gosu putiuser puti
fi

exec "$@"