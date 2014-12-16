NAME = moul/go-dl-extract
VERSIONS = Darwin-x86_64 Linux-armel Linux-armhf Linux-i386 Linux-x86_64
DEFAULT_VERSION = Linux-x86_64

all:	build
	docker tag $(NAME):$(DEFAULT_VERSION) $(NAME):latest

build:	generate
	for version in $(VERSIONS); do \
	    cd $$version; \
	    docker build -t $(NAME):$$version .; \
	    cd ..; \
	done

generate:
	for version in $(VERSIONS); do \
	    mkdir -p $$version; \
	    cp ../dist/go-dl-extract-$$version ../dist/ssl.tar $$version/; \
	    echo "FROM scratch" > $$version/Dockerfile; \
	    echo "MAINTAINER Manfred Touron <m@42.am> (@moul)" >> $$version/Dockerfile; \
	    echo "ADD ssl.tar /" >> $$version/Dockerfile; \
	    echo "ADD go-dl-extract-$$version /bin/sh" >> $$version/Dockerfile; \
	done
