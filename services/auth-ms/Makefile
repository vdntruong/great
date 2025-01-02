ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

TEMPORARY_DIR := tmp
CONTAINER_NAME := auth-ms
DOCKER_IMAGE_NAME := theg-auth-ms

.PHONY: lint
lint:
	golangci-lint run ./... --verbose

.PHONY: setup-tmp
setup-tmp:
	mkdir $(TEMPORARY_DIR) || true
	rsync -av --progress ../commons $(TEMPORARY_DIR)
	rsync -av --progress ../auth-ms $(TEMPORARY_DIR) --exclude auth-ms/deployments
	cp ../go.work.sum $(TEMPORARY_DIR)/

.PHONY: clean-tmp
clean-tmp:
	rm -rf $(TEMPORARY_DIR)

.PHONY: docker-image
docker-image:
	docker build -t $(DOCKER_IMAGE_NAME) -f deployments/Dockerfile $(TEMPORARY_DIR)/

.PHONY: build
build: clean-tmp setup-tmp docker-image clean-tmp

.PHONY: start-container
start-container:
	docker run -it --rm -p 8081:8080 --name $(CONTAINER_NAME) $(DOCKER_IMAGE_NAME)
