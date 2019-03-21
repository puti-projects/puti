#!/bin/bash
set -e

# Create VOLUME folder if not exist
for f in /data/configs /data/theme /data/uploads; do
    if ! test -d $f; then
        echo "$f is not exist, create it"
        mkdir -p $f
    fi
done

# Only for the first time
if ! test -d /app/puti/configs/config.yaml; then
    cp /app/puti/configs/config.yaml.example /app/puti/configs/config.yaml
    # set owner
	chown -R putiuser /data /app/puti
fi
chmod 0755 /data /data/configs /data/theme /data/uploads

# Link volumed data with app data
cd /app/puti
ln -sfn /data/configs  ./configs
ln -sfn /data/theme  ./theme
ln -sfn /data/uploads  ./uploads
