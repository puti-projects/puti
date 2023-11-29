# Changelog

## v0.4.2 Go Version Update [2023-11-28]

### 功能和优化
- 知识库前端支持局部加载

### Bug 修复


---

## v0.4.1 Fixed  [2021-03-23]

### 功能和优化
- 优化 Auto HTTPS 和 Health Check；目前不做重定向判断，如根域名到 www 域名，HTTP 到 HTTPS；较好的实践是再加一层 WebServer。
- 主题支持 favicon.ico 自定义

### Bug 修复
- 修复一个缓存刷新问题

---

## v0.4.0 Knowledge Base [2020-11-26]

### 功能和优化
- 新增知识库大板块（含笔记和文档；文档有多版本功能而笔记没有）
- 前台部分页面增加缓存

---

## v0.3.1 Auto HTTPS [2020-11-06]

### 功能和优化
- 目录优化、CI 优化，代码质量检查
- 增加自动 HTTPS（autocert），优化配置
- 简化部署
- 移除 Circle-CI，Github Action 取代

### Bug 修复
- 修复 GORM v2 带来的兼容性问题
  
---

## v0.3.0 Optimization [2020-09-12]

### 功能和优化
- 优化 Lin 主题 Latex 公式样式在移动端的展示
- 移除 vendor 目录和依赖，使用 goproxy
- 支持 Github Actions
- 大量基础组件优化和重构
- 后台 api 结构重构
- 升级 GROM 到 v2

### Bug 修复
- 修复标签修改出错的问题

---

## v0.2.2 Article Cover Picture [2019-08-14]

### 功能和优化
- 新增或编辑文章时现在可以选择封面图了
- 优化后台管理媒体库展示
- 优化 Lin 主题样式

### Bug 修复
- 修复加载主题时加载错误的文件（非文件夹）

---

## v0.2.1 Bug fixed and small adjustment [2019-08-06]

### 功能和优化
- 调整后台管理样式
- 优化 Lin 主题的一些样式

### Bug 修复
- 修复创建或编辑文章时，选择专题的树形结构取值不正确导致文章数计算不对
- 修复 Lin 主题页面和归档页侧栏无法展开
- 修复仪表盘系统信息 widget 的显示错误

---

## v0.2.0 New Theme [2019-07-21]

### 功能和优化

- 新增新主题 Lin
- 新增主题管理，可以自由切换不同的主题
- 优化 logger 日志组件

### Bug 修复

- 修复创建或者编辑文章时，选择分类的树形强关联，导致提交的数据不正确
- 修复 Emma 主题 katex css 引入的 bug

---

## v0.1.0 Fisrt Version [2019-04-01]

第一个版本，包含基本功能，已经初步成型。包含以下内容：

- 登录
- 文章
- 页面
- 分类
- 标签
- 专题
- 媒体
- 用户
- 普通设置
- 主题支持
- HTTPS
- WebServer 转发
- Docker 镜像支持
- 等
