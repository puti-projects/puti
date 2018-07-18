# gingob
Golang+Vue

# Build by Docker
容器内编译应用：
```
$ docker run --rm -v "$PWD":/go/src/gingob -w /go/src/gingob golang go build -v
```
会存在依赖包不存在的问题，所以把src整个挂载：
```
$ docker run --rm -v E:/GoPath/src:/go/src -w /go/src/gingob golang go build -v
```
交叉编译，例如编译一个Windows平台二进制文件:
```
$ docker run --rm -v E:/GoPath/src:/go/src -w /go/src/gingob -e GOOS=windows -e GOARCH=386 golang go build -v
```
容器内通过 Makefile 编译应用：
```
$ docker run --rm -v "$PWD":/go/src/gingob -w /go/src/gingob golang make
```

如果要在容器内运行应用，则编写 dockerfile：
```
FROM golang

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
```
然后运行：
```
$ docker build -t my-golang-app .
$ docker run -it --rm --name my-running-app my-golang-app
```
