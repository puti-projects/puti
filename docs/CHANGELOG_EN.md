# Changelog

## v0.4.2 Go Version Upgrade [2023-11-30]

# Feature and Enhancement
- Go version updated to 1.21
- Front-end of the knowledge base supports partial loading
- Automatic HTTPS is enabled by default
- Default redirects are added, from non-www to www, and from http to https
- DockerHub image push is replaced by Github Action
- Jwt package is switched to [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)

# Bug fixes
- Failure in gosu installation in dockerfile

---

## v0.4.1 Fixed  [2021-03-23]

### Feature and Enhancement
- Optimize Auto HTTPS and Health Check; currently, no redirection judgment is made, such as root domain name to www domain name, HTTP to HTTPS; a better practice is to add another layer of WebServer.
- Theme supports a custom favicon.ico file now

### Bug Fixed
- Fix a cache refresh problem

---

## v0.4.0 Knowledge Base [2020-11-26]

### Feature and Enhancement
- Added knowledge base section (including notes and documents; documents have multi-version functions but notes do not)
- Increase cache for some pages in the foreground

---

## v0.3.1 Auto HTTPS [2020-11-06]

### Feature and Enhancement
- Catalog optimization, CI optimization, code quality inspection
- Add automatic HTTPS (autocert), optimize configuration
- Simplified deployment
- Remove Circle-CI and replace it with Github Action

### Bug Fixed
- Fix compatibility issues brought by GORM v2
  
---

## v0.3.0 Optimization [2020-09-12]

### Feature and Enhancement
- Optimize the display of the Lin theme Latex formula style on the mobile terminal
- Remove vendor directory and dependencies, use goproxy
- Support Github Actions
- A large number of basic component optimization and reconstruction
- Back-end api structure reconstruction
- Upgrade GROM to v2

### Bug Fixed
- Fix the problem of incorrect label modification

---

## v0.2.2 Article Cover Picture [2019-08-14]

### Feature and Enhancement
- You can now choose a cover image when you add or edit an article now.
- Optimize background management media library style display
- Optimize theme Lin's styles

### Bug Fixed
- Fixed loading wrong files (non-folder) when loading theme

---

## v0.2.1 Bug fixed and small adjustment [2019-08-06]

### Feature and Enhancement
- Adjusted background management styles
- Optimized some styles of theme Lin

### Bug Fixed
- Fixed the tree structure of subject which cause the number of articles is incorrectly calculate when created or edited article
- Fixed the theme Lin's page and archive page sidebar cannot be expanded
- Fixed display error of dashboard system information widget

---

## v0.2.0 New Theme [2019-07-21]

### Feature and Enhancement

- Added new theme Lin
- Added theme management, now we can switch installed theme simply
- Optimized logger component

### Bug Fixed

- Fixed catergory select to strectly tree when creating or editing article.
- Fixed theme Emma import katex css in wrong way

---

## v0.1.0 Fisrt Version [2019-04-01]

The first version, which contains basic features, has been initially formed. Contains the following:

- Log in
- Article
- Page
- Category
- Tag
- Subject
- Media
- User
- Normal settings
- Theme support
- HTTPS
- WebServer forwarding
- Docker image support
- etc.
