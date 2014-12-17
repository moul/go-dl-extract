go-dl-extract
=============

A static binary that mimics the `ADD and extract` Dockerfile's command for
remote tarballs.

Usage
-----

    FROM moul/go-dl-extract
    MAINTAINER Manfred Touron <m@42.am> (@moul)
    # by inheriting the moul/go-dl-extract, the first RUN means a remote ADD
    RUN http://archlinuxarm.org/os/ArchLinuxARM-armv7-latest.tar.gz
    CMD ["/bin/sh"]

Context
-------

The way to create a distribution/base image is to `ADD` a local tarball to a
`FROM scratch` Dockerfile.

The way to create a **trusted** ditribution image is to put the tarball in the
Github repository with the Dockerfile.
Github prevents the files >100MB to be uploaded.

During a short period, the `ADD` command was also extracting remote tarballs,
but it changed with [this PR](https://github.com/docker/docker/pull/4193).

Security note
-------------

The main advantage to keep the tarball in the Github repository is to create a
trusted build that will always give the same result and also provides the 
ability to be analyzed.
([example](https://github.com/tianon/docker-brew-ubuntu-core/tree/dist/utopic))

`go-dl-extract` will print the tarball checksum and can also do a comparison
with `--md5=THECHECKSUM`.

If the code is using a `-latest.tar.gz` kind of tarball, it is recommended to
mirror the tarball somewhere and tag the tarball with the today date.

Cross platform
--------------

This project was built using the `golang:1.3-cross` image and was cross-compiled
for 5 systems and architectures (Docker images availables):

System | Architecture | Docker image size | Comment                 | Docker image
-------|--------------|-------------------|-------------------------|----------------------------------
Darwin | x86_64       | 5.8MB             | Works with boot2docker  | moul/go-dl-extract:Darwin-x86_64
Linux  | x86_64       | 5.4MB             | Also `latest` tag       | moul/go-dl-extract:Linux-x86_64
Linux  | i386         | 4.4MB             |                         | moul/go-dl-extract:Linux-i386
Linux  | armel        | 4.4MB             |                         | moul/go-dl-extract:Linux-armel
Linux  | armhf        | 4.4MB             | Works on Online-Labs C1 | moul/go-dl-extract:Linux-armhf

The compiled binaries are also available in the
[dist](https://github.com/moul/go-dl-extract/tree/dist/dist) branch.

Dependents
----------

- trusted [armbuild/archlinux-disk](https://registry.hub.docker.com/u/armbuild/archlinux-disk/dockerfile/) image

Related discussions on Docker
-----------------------------

- https://github.com/docker/docker/pull/4193
- https://github.com/docker/docker/issues/3050
- https://groups.google.com/forum/#!topic/docker-user/0aSY7R59qqo
- https://github.com/docker/docker/issues/3964
