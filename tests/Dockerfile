FROM scratch
ADD ssl.tar /
ADD go-dl-extract-Linux-x86_64 /bin/sh
RUN --md5=5f9cbd2728532c16ac28a501b3a59b7e https://github.com/jpetazzo/docker-busybox/raw/buildroot-2014.02/tarmaker-buildroot/rootfs.tar

RUN echo "Hello World"
