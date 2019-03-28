# Docker for Puti
<a href="./README.md">Engilsh</a> | 中文

## Usage

```sh
# 从 Docker Hub 拉取镜像
$ docker pull puti/puti

# 创建需要挂载的目录，例如：`/data/puti`为应用文件存放目录，`/data/logs/puti`为日志存放目录
$ mkdir -p /data/puti /data/logs/puti

# 第一次通过`docker run`来创建一个容器
$ docker run --name=puti -p 80:8000 -p 443:8080 -v /data/puti:/data/puti -v /data/logs/puti:/data/logs/puti puti/puti

# 使用 `docker stop``docker start`来停止，关闭容器。
$ docker stop puti
$ docker start puti
```

## VOLUME
通过以上例子创建并运行之后，我们会得到以下结构的文件：

    /data
    |-- puti
        |-- configs
            |-- config.yaml # 配置文件
            |-- config.yaml.example
        |-- theme
            |-- # 主题文件和目录
        |-- uploads
            |-- # 上传的图片等文件和目录
    |-- logs
        |-- puti
            |-- server.log # Puti 应用日志文件

configs 下的 config.yaml 为生成的 Puti 应用配置文件，需要修改对应的配置，比如HTTP端口，数据库连接等。
