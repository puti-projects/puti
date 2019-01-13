BASEDIR = $(shell pwd)

# build with verison infos
versionDir = "github.com/puti-projects/puti/internal/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: build
build:
	@echo "Building binary file."
	go build -mod=vendor -v -ldflags ${ldflags} .
clean:
	@echo "Cleaning."
	rm -f puti
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	@echo "Running go tool."
	go tool vet .
test:
	@echo "Testing."
	go test -v ./...
ca:
	@echo "Generating ca files."
	openssl req -new -nodes -x509 -out configs/server.crt -keyout configs/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help:
	@echo "make - run gotool and build"
	@echo "make build - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'vet'"
	@echo "make test - run go test"
	@echo "make ca - generate ca files"

.PHONY: build clean test gotool ca help
