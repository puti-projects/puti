# Docker for Puti
Engilsh | <a href="./README_CN.md">中文</a>

## Usage

```sh
# Pull image from Docker Hub.
$ docker pull puti/puti

# Create local directory for volume.
$ mkdir -p /data/puti /data/logs/puti

# Use `docker run` for the first time.
$ docker run --name=puti -p 80:8000 -p 443:8080 -v /data/puti:/data/puti -v /data/logs/puti:/data/logs/puti puti/puti

# Use `docker start` if you have stopped it.
$ docker stop puti
$ docker start puti
```

## VOLUME
In this case, application files will be store in the local path `/data/puti`, and log files  will be store in path `/data/logs/puti`
The structure will be like：

    /data
    |-- puti
        |-- configs
            |-- config.yaml # config file
            |-- config.yaml.example
        |-- theme
            |-- # Theme paths and files
        |-- uploads
            |-- # Uploaded files
    |-- logs
        |-- puti
            |-- server.log # Puti application log file

The config.yaml under the configs path is the config file of Puti appliction, you should change some configuration and restart the container such as http ports, database connection etc.
