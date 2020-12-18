# Golang docker scratch

Use docker build-kit to speed up containerized golang build.

## Usage

```bash
make docker

# custom image tag
make docker IMAGE_TAG="go-test"
docker run --rm go-test

# use make run
make run IMAGE_TAG="go-test"
```

## GO Proxy

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