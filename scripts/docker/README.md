# Docker for Puti

## Usage

```sh
# Pull image from Docker Hub.
$ docker pull puti/puti

# Create local directory for volume.
$ mkdir -p /data/puti

# Use `docker run` for the first time.
$ docker run --name=puti -p 8000:8000 -p 8080:8080 -v /data/puti:/data puti/puti

# Use `docker start` if you have stopped it.
$ docker start puti
```