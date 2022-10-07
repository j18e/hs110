NAME := hs110-exporter
COMMIT_HASH := $(shell git rev-parse --short HEAD)
IMAGE_NAME := j18e/$(NAME)
IMAGE_FULL := $(IMAGE_NAME):$(COMMIT_HASH)

build:
	GOOS=linux GOARCH=amd64 go build -trimpath -o ./$(NAME) ./cmd/$(NAME)

docker-build:
	docker build -t $(IMAGE_FULL) .

docker-push:
	docker push $(IMAGE_FULL)

all: build docker-build docker-push
