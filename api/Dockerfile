FROM golang

RUN go get github.com/canthefason/go-watcher/cmd/watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher
WORKDIR /go/src/github.com/sonofbytes/gocrawl

CMD ["/go/bin/watcher", "-run", "github.com/sonofbytes/gocrawl/api"]
EXPOSE 8888