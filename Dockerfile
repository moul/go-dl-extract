FROM golang:1.3-cross

ENV CGO_ENABLED 0
ADD go-dl-extract.go /go/

RUN GOOS=darwin GOARCH=amd64          go build -v -ldflags -d -o /go/bin/go-dl-extract-osx
RUN GOOS=linux  GOARCH=amd64          go build -v -ldflags -d -o /go/bin/go-dl-extract-amd64
RUN GOOS=linux  GOARCH=386            go build -v -ldflags -d -o /go/bin/go-dl-extract-i386
RUN GOOS=linux  GOARCH=arm   GOARM=5  go build -v -ldflags -d -o /go/bin/go-dl-extract-armel
RUN GOOS=linux  GOARCH=arm   GOARM=6  go build -v -ldflags -d -o /go/bin/go-dl-extract-armhf
