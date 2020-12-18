# Golang docker scratch

[![Build](https://github.com/lingsamuel/golang-docker-scratch/workflows/build/badge.svg)](https://github.com/lingsamuel/golang-docker-scratch/actions?query=workflow%3Abuild)

Use docker build-kit to speed up containerized golang build.

Comes from: [jeremyhuiskamp/golang-docker-scratch](https://github.com/jeremyhuiskamp/golang-docker-scratch)

This is a `Template` repo, you can click the `Use this template` button to generate new repo with same files.

## Usage

```bash
make docker

# custom image tag
make docker IMAGE_TAG="go-test"
docker run --rm go-test

# use make run
make run IMAGE_TAG="go-test"
```

## GOPROXY

Change `ENV GOPROXY=https://goproxy.cn` in Dockerfile.

## Without `make`

Replace:

```dockerfile
RUN --mount=type=cache,target=/root/.cache/go-build make build
```

to:

```dockerfile
RUN --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o output/main ./cmd/root.go
```

Then build with:

```bash
DOCKER_BUILDKIT=1 docker build -f ./Dockerfile -t "go-scratch" .
```

## Special Notes

You may need docker build later than 2019-04-22 if you are using registry (see [moby/buildkit#779](https://github.com/moby/buildkit/issues/779)).
