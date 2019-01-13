# Puti
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fputi-projects%2Fputi.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fputi-projects%2Fputi?ref=badge_shield)   


[中文](https://github.com/puti-projects/puti/blob/master/README.md)
 | Engilsh

------------


## Project Status
This project is still developing.

## Features
* [x] Loign
* [ ] Register
* [x] Article
* [x] Category
* [x] Tag
* [ ] Subject
* [ ] Link
* [x] Media
* [x] Page
* [x] User
* [ ] Comments
* [x] Setting
* [x] Theme support
* [ ] Plugin support

## Screenshot

## Online Demo
[Demo](https://demo.goozp.com)

## Environmental requirements
- Golang 1.11 (Build using modules)
- MySQL
- Nginx (Optional)

## Getting Started

### Configuration
### Installation
### Build by Docker
- Using a MySQL in lcoal   

运行一个mysql容器供本地使用
```
docker run --name go-mysql -p 3306:3306 -v E:/data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql
```

- Compile the Puti app inside the Docker container 

容器内编译应用：
```
$ docker run --rm -v "$PWD":/go/src/puti -w /go/src/puti golang go build -v
```
会存在依赖包不存在的问题，所以把src整个挂载：
```
$ docker run --rm -v E:/GoPath/src:/go/src -w /go/src/puti golang go build -v
```

- Cross-compile your app inside the Docker container

交叉编译，例如编译一个Windows平台二进制文件:
```
$ docker run --rm -v E:/GoPath/src:/go/src -w /go/src/puti -e GOOS=windows -e GOARCH=amd64 golang go build -v
```

- Build by a Makefile (recommend)   

容器内通过 Makefile 编译应用：
```
$ docker run --rm -v "$PWD":/go/src/puti -w /go/src/puti golang make
$ docker run --rm -v E:/GoPath/src:/go/src -w /go/src/puti golang make
$ docker run --rm -v E:/GoPath/src:/go/src -w /go/src/puti -e GOOS=windows -e GOARCH=amd64 golang make
```

- Start a Go instance in your app   
如果要在容器内运行应用，通过运行 dockerfile 构建镜像来使用：
```
$ docker build -f build/package/Dockerfile -t puti .
$ docker run -it --rm --name puti puti
```

### Usage

## Theme
More themes is creating.

## Documentation
TODO

## Changelog
Detailed changes for each release are documented in the [changelog file]((https://github.com/axetroy/vscode-gpm/blob/master/CHANGELOG.md)).

## Dependencies
Thanks for these great open source libraries:
| Dependency | About |
| :----- | :----- | 
| [gin-gonic/gin](https://github.com/gin-gonic/gin) |  HTTP web framework written in Go |
| [jinzhu/gorm](https://github.com/jinzhu/gorm) | The ORM library for Golang|
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