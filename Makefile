SHELL = /bin/bash
IMAGE_TAG := go-scratch

build:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o output/main ./cmd/root.go
clean:
	rm -rf output
docker:
	DOCKER_BUILDKIT=1 docker build -f ./Dockerfile -t $(IMAGE_TAG) .
run: build docker
	docker run $(IMAGE_TAG)
