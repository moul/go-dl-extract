FROM scratch
ADD ssl.tar /
ADD go-dl-extract-Linux-x86_64 /bin/sh
RUN https://github.com/jpetazzo/docker-busybox/raw/buildroot-2014.02/tarmaker-buildroot/rootfs.tar

RUN echo "Hello World"
