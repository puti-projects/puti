FROM golang
LABEL maintainer="gzp@goozp.com"

WORKDIR /go/src/gingob
COPY . .

# go get 包含许多golang.org包，会失败
RUN go get -d -v ./...
RUN make

CMD ["gingob"]