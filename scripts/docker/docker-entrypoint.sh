#!/bin/bash
set -e

# Create VOLUME folder if not exist
for f in /data/configs /data/theme /data/uploads; do
    if ! test -d $f; then
        echo "$f is not exist, create"
        mkdir -p $f
    fi
done

cd /app/puti

# Link volumed data with app data
ln -sfn /data/configs  ./configs
ln -sfn /data/theme  ./theme
ln -sfn /data/uploads  ./uploadss

# Only for the first time
if ! test -d /data/configs/config.yaml; then
    # set owner
	chown -R putiuser:putiuser /data /app/puti
fi
chmod 0755 /data /data/configs /data/theme /data/uploads
