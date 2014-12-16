FROM golang:1.3-cross

ENV CGO_ENABLED 0

# Install Godep for vendoring
RUN go get github.com/tools/godep
# Recompile the standard library without CGO
RUN go install -a std

# Declare the maintainer
MAINTAINER Manfred Touron @moul

# For convenience, set an env variable with the path of the code
ENV APP_DIR /go

ADD . /go/

# Compile the binary and statically link
RUN cd $APP_DIR && GOOS=darwin GOARCH=amd64          godep go build -a -v -ldflags '-d -w -s' -o /go/bin/go-dl-extract-Darwin-x86_64
RUN cd $APP_DIR && GOOS=linux  GOARCH=amd64          godep go build -a -v -ldflags '-d -w -s' -o /go/bin/go-dl-extract-Linux-x86_64
RUN cd $APP_DIR && GOOS=linux  GOARCH=386            godep go build -a -v -ldflags '-d -w -s' -o /go/bin/go-dl-extract-Linux-i386
RUN cd $APP_DIR && GOOS=linux  GOARCH=arm   GOARM=5  godep go build -a -v -ldflags '-d -w -s' -o /go/bin/go-dl-extract-Linux-armel
RUN cd $APP_DIR && GOOS=linux  GOARCH=arm   GOARM=6  godep go build -a -v -ldflags '-d -w -s' -o /go/bin/go-dl-extract-Linux-armhf
