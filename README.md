# tinyvm-go

TinyVM-go is a go implement version that inspired by the project [![tinyvm](https://github.com/jakogut/tinyvm)](https://github.com/jakogut/tinyvm)

TinyVM is a virtual machine with the goal of having a small footprint.
Translating the source code into bytecodes that we can operate.

## Building

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all architectures.

Run `make push` to push the container image to `REGISTRY`.  Run `make all-push`
to push the container images for all architectures.

Run `make clean` to clean up.
