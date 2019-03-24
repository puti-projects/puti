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
* [ ] Toc support(not theme)

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
#### Using Docker
##### 使用现成的镜像
```sh
# Pull image from Docker Hub.
$ docker pull puti/puti

# Create local directory for volume.
$ mkdir -p /data/puti

# Use `docker run` for the first time.
$ docker run --name=puti -p 80:8000 -p 443:8080 -v /data/puti:/data puti/puti

# Use `docker start` if you have stopped it.
$ docker start puti
```  

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