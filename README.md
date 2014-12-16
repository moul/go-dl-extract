go-dl-extract
=============

The goal of this project is to create a minimal image that support and `ADD` and `extract` of a remote tarball.

Context
-------

The way to create a trusted build of distribution is to put the tarball in a Github repos and a Dockerfile that ADDs the tarball locally.

But Github blocks the files with a length > 100Mb

Since [this PR](https://github.com/docker/docker/pull/4193), we cannot ADD and untar a tarball using the URL

Related links
-------------

- https://github.com/docker/docker/pull/4193
- https://github.com/docker/docker/issues/3050
- https://groups.google.com/forum/#!topic/docker-user/0aSY7R59qqo
- https://github.com/docker/docker/issues/3964
