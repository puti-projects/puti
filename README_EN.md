<p align="center">
    <img src="assets/logo.png" alt="Puti Logo" width="150" height="150">
</p>
<h1 align="center">Puti</h1>
<p align="center">
    <em>:black_nib: Puti is a full-featured writing system written in Golang.</em>
</p>
<p align="center">
    <a href="https://github.com/puti-projects/puti/actions?query=workflow%3Abuild">
        <img src="https://github.com/puti-projects/puti/workflows/build/badge.svg" alt="Github Actions">
    </a>
    <a href="https://codecov.io/gh/puti-projects/puti">
        <img src="https://codecov.io/gh/puti-projects/puti/branch/master/graph/badge.svg" />
    </a>
    <a href="https://goreportcard.com/report/github.com/puti-projects/puti">
        <img src="https://goreportcard.com/badge/github.com/puti-projects/puti" alt="Go Report Card">
    </a>
    <a href="https://github.com/puti-projects/puti/releases">
        <img src="https://img.shields.io/github/release/puti-projects/puti.svg?style=flat" alt="Release">
    </a>
    <a href="https://github.com/puti-projects/puti/blob/master/LICENSE">
        <img src="https://img.shields.io/badge/License-GPLv3.0-important.svg?style=flat" alt="license" />
    </a>
</p>
<p align="center">
<a href="https://github.com/puti-projects/puti/blob/master/README.md">中文</a>
 | Engilsh
</p>

## Project Status

This project is still developing, and the goal is the next generation writing system for geeks.

## Environmental requirements

- Golang 1.13+ (Build using modules)
- MySQL
  
This project uses Go Modules, so it is recommended to use Go 1.13 or above; The project does not rely on Web Server such as Nginx, but you can configure and use Nginx.

## Features

The project plan implementation and the functions that have been implemented are as follows:

* [ ] Features
  * [x] User
  * [ ] Loign and register
    * [x] Loign
    * [ ] Register
    * [ ] Third party access (github, etc.)
  * [x] Blog system
    * [x] Article
    * [x] Page
    * [x] Category
    * [x] Tag
    * [x] Subject
  * [x] Knowledge base system
    * [x] Notebook
    * [x] Documentation set
  * [x] Media
  * [ ] Link
  * [ ] Comments
  * [ ] Settings
    * [x] Normal setting
    * [ ] Third party settings (GitHub, WeChat, etc.)
  * [ ] Theme
    * [X] Theme support
    * [X] Default theme (Emma and Lin)
    * [X] Free switching
    * [ ] Theme template files modification(backend console)
  * [ ] Plugin
    * [ ] Plugin support
    * [ ] Plugin management (upload, delete, etc.)
  * [ ] Email
    * [ ] Mail configuration
    * [ ] Mail delivery
* [ ] Technical Support
  * [ ] Complete i18n support
  * [ ] Mail service configuration
  * [ ] Toc support (not theme)
  * [ ] Configure image cropping
  * [X] HTTPS (Support automatic HTTPS)
  * [x] WebService forwarding
  * [ ] Avatar access
  * [ ] OAuth 
  * [ ] Media file cloud storage (for CDN)
* [ ] Ecology
  * [x] Docker image support
  * [x] Configured automatic deployment script  
  * [ ] Simple statistical system

## Screenshot

![screenshot_1](./docs/images/screenshot_one.png)
![screenshot_2](./docs/images/screenshot_two.png)
![screenshot_3](./docs/images/screenshot_three.png)

## Getting Started

### Configuration

Puti's configuration file is `config.yaml` under path `configs`, and the configuration file can be initialized from `config.yaml.example` when first used.  
Configuration to be aware of：

| Configuration | Description |
| :----- | :----- |
| server.http_port | HTTP Port |
| server.https_open |  Open HTTPS  |
| server.auto_cert |  Open auto cert  |
| server.https_port |  HTTPS Port  |
| server.tls_cert | If it is not automatic cert，the SSL certificate path   |
| server.tls_key |  If it is not automatic cert，the SSL private key path  |
| safety.jwt_secret |  Json web token secret key |
| db.name |  Database name  |
| db.addr |  Database HOST:PORT  |
| db.username |  Database user  |
| db.password |  Database password |

### Installation

#### Source installation

The project uses Go Module, so Go 1.13 and above are required. The Vendor directory is currently removed, because now `go proxy` can solve some problems well.

```sh
# Download
$ go get -u github.com/puti-projects/puti

# Use Makefile to build programs
$ cd $GOPATH/src/github.com/puti-projects/puti
$ make
```

#### Using Docker

##### Using Ready-made Docker Images

We have provided a ready-made image that can be pulled directly:

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

More information：[Docker use](./scripts/docker.README.md)

##### Use configurable deployment script

We provide a one-click deployment of the Docker-compose script file, which is convenience for build the working environment. [puti-projects/puti-environment](https://github.com/puti-projects/puti-environment)

### Usage
If initialization failed, which may be a problem with the database configuration (currently there is no installation guide). An account is initialized by default with the default account `admin` and password `admin`. Please create your own account and remove the default account. Installation and guidance will be considered after the functions are complete.

## Theme

More themes is creating. Now we have two different style default themes, Emma and Lin.

## Documentation

TODO

## Changelog

Detailed changes for each release are documented in the [changelog file]((https://github.com/axetroy/vscode-gpm/blob/master/CHANGELOG.md)).

## Dependencies

Thanks for these great open source libraries:

| Dependency | About |
| :----- | :----- | 
| [gin-gonic/gin](https://github.com/gin-gonic/gin) |  HTTP web framework written in Go. |
| [go-gorm/gorm](https://github.com/go-gorm/gorm)  | The ORM library for Golang. |
| [allegro/bigcache](https://github.com/allegro/bigcache) | Efficient cache for gigabytes of data written in Go. |
| [spf13/viper](https://github.com/spf13/viper) |  Complete configuration solution. |
| [go.uber.org/zap](https://go.uber.org/zap) |  Fast, structured, leveled logging. |
| [vuejs/vue](https://github.com/vuejs/vue) | JavaScript framework for building UI on the web. |
| [ElemeFE/element](https://github.com/ElemeFE/element) | A Vue.js 2.0 UI Toolkit for Web.  |
| [PanJiaChen/vue-element-admin](https://github.com/PanJiaChen/vue-element-admin) | A front-end management background integration solution. |
| [hinesboy/mavonEditor](https://github.com/hinesboy/mavonEditor) (will be removed) | A markdown editor. |
| [Vanessa219/vditor](https://github.com/Vanessa219/vditor) | An in-browser markdown editor. |
| [88250/lute](https://github.com/88250/lute) | A structured Markdown engine that supports Go and JavaScript. |

## Description

### Deploy
It is not necessary to use a WebServer such as Nginx, and supports automatic HTTPS; currently, no redirection judgment is made, such as root domain to `www` domain, HTTP to HTTPS; better practice is to add another layer of WebServer.


## Contributors
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
| [<img src="https://avatars3.githubusercontent.com/u/17734933?s=460&v=4" width="100px;"/><br /><sub>goozp</sub>](https://www.goozp.com)<br />[💻](https://github.com/puti-projects/puti/commits?author=goozp "Code commitor")[📚](https://github.com/dawnlabs/carbon/commits?author=briandennis "Documentation")[🎨](#design "Design") |
| :---: |

<!-- ALL-CONTRIBUTORS-LIST:END -->

## Thanks
Thanks to JetBrains for providing free Goland IDE based on JetBrains OS licenses.

[![goland](./docs/images/icon-goland.svg)](https://www.jetbrains.com/?from=puti)


## License 

Puti is under the GPL-3.0 license. See the [LICENSE](https://github.com/puti-projects/puti/blob/master/LICENSE) file for details.