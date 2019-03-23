#!/bin/sh
set -e

#Create VOLUME folder if not exist
BASE_DATA=/data/puti
BASE_APPPUTI=/app/puti
for f in /configs /theme /uploads; do
    if ! test -d ${BASE_DATA}$f; then
        echo "${BASE_DATA}$f is not exist, move from ${BASE_APPPUTI}$f"
        mv ${BASE_APPPUTI}$f ${BASE_DATA}$f
    fi
done

if ! test -d /data/logs/puti; then
    echo "/data/logs/puti is not exist, creating."
    mkdir -p /data/logs/puti
fi

chmod 0755 /data/puti/configs /data/puti/theme /data/puti/uploads /data/logs/puti

# Link volumed data with app data
for f in /configs /theme /uploads; do
    if ! test -L ${BASE_APPPUTI}$f; then
        echo "link ${BASE_APPPUTI}$f is not exist, creating from source: ${BASE_DATA}$f"
        ln -sfn ${BASE_DATA}$f ${BASE_APPPUTI}$f
    fi
done

# Only for the first time
if ! test -d /app/puti/configs/config.yaml; then
    cp /app/puti/configs/config.yaml.example /app/puti/configs/config.yaml
fi

exec "$@"
