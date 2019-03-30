<p align="center">
    <img src="assets/logo.png" alt="Puti Logo" width="150" height="150">
</p>
<h1 align="center">Puti</h1>
<p align="center">
    <em>:black_nib: Puti is a blog system written in Golang.</em>
</p>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/puti-projects/puti">
        <img src="https://goreportcard.com/badge/github.com/puti-projects/puti" alt="Go Report Card">
    </a>
    <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Fputi-projects%2Fputi?ref=badge_shield">
        <img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Fputi-projects%2Fputi.svg?type=shield" alt="FOSSA Status">
    </a>
</p>
<p align="center">
中文
 | <a href="https://github.com/puti-projects/puti/blob/master/docs/README_EN.md">Engilsh</a>
</p>


## Project Status
Puti 项目现在仍在开发中。因为是作者接触 Go 语言的第一个项目，所以代码质量不到位之处，将会在未来不断优化，非常欢迎你的贡献。

## Environmental requirements
 - Golang 1.11+ (Build using modules)
 - MySQL 
 - Nginx (Optional)
  
Golang 1.11 版本开始支持 go module，本项目使用了go module；Nginx 为可选配置。

## Features
项目计划实现以及已经实现的功能如下：
* [ ] 功能
  * [ ] 登录注册
    * [x] 登录
    * [ ] 注册
    * [ ] 第三方接入（github等）
  * [x] 文章
  * [x] 页面
  * [x] 分类
  * [x] 标签
  * [x] 专题
  * [ ] 链接
  * [x] 媒体
  * [x] 用户
  * [ ] 评论
  * [ ] 设置
    * [x] 普通设置
    * [ ] 第三方设置（接入GItHub，WeChat等）
  * [ ] 前台主题
    * [X] 主题支持
    * [X] 默认主题（Emma）
    * [ ] 自由切换 
  * [ ] 插件
    * [ ] 插件支持
    * [ ] 插件管理（上传、删除等）
  * [ ] 邮件
    * [ ] 邮件配置
    * [ ] 邮件发送
* [ ] 技术支持 
  * [ ] 完善的 i18n 
  * [ ] 邮件服务配置
  * [ ] TOC (目前在前端主题实现)
  * [ ] 配置图片裁切
  * [X] HTTPS
  * [x] WebServer 转发
  * [ ] 头像接入
  * [ ] OAuth 
  * [ ] 媒体文件云存储
* [ ] 生态
  * [x] Docker 镜像支持
  * [x] 配置化的自动部署脚本  
  * [ ] 简单的统计系统

## Screenshot
![Docker use](./docs/images/screenshot1.png)
![Docker use](./docs/images/screenshot2.png)

## Getting Started

### Configuration
### Installation
#### Using Docker
##### 使用现成的镜像
我们已经提供了现成的镜像，可以直接拉取使用：
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
更多内容查看：[Docker use](./script/docker/README.md)

##### 使用可配置的部署脚本
我们提供了简单方便地一键部署 Docker-compose 脚本文件，懒人必备。具体使用查看：[puti-projects/puti-environment](https://github.com/puti-projects/puti-environment)

### Usage

## Theme
More themes is creating.

## Documentation
TODO

## Changelog
Detailed changes for each release are documented in the [changelog file]((https://github.com/axetroy/vscode-gpm/blob/master/CHANGELOG.md)).

## Dependencies
| Dependency | About |
| :----- | :----- | 
| [gin-gonic/gin](https://github.com/gin-gonic/gin) |  HTTP web framework written in Go |
| [jinzhu/gorm](https://github.com/jinzhu/gorm) | The ORM library for Golang|
| [patrickmn/go-cache](https://github.com/patrickmn/go-cache) | An in-memory key:value store/cache|
| [spf13/viper](https://github.com/spf13/viper) |  complete configuration solution|
| [go.uber.org/zap](https://go.uber.org/zap) |  fast, structured, leveled logging|
| [vuejs/vue](https://github.com/vuejs/vue) | JavaScript framework for building UI on the web |
| [ElemeFE/element](https://github.com/ElemeFE/element) | A Vue.js 2.0 UI Toolkit for Web  |
| [PanJiaChen/vue-element-admin](https://github.com/PanJiaChen/vue-element-admin) | A front-end management background integration solution |

## Q & A


## Authors
Puti is a project by 
- Goozp ([@goozp](https://www.goozp.com))

## Contributors
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
| [<img src="https://avatars3.githubusercontent.com/u/17734933?s=460&v=4" width="100px;"/><br /><sub>goozp</sub>](https://www.goozp.com)<br />[💻](https://github.com/puti-projects/puti/commits?author=goozp "Code commitor")[📚](https://github.com/dawnlabs/carbon/commits?author=briandennis "Documentation")[🎨](#design "Design") | 
| :---: |

<!-- ALL-CONTRIBUTORS-LIST:END -->

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fputi-projects%2Fputi.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fputi-projects%2Fputi?ref=badge_large)