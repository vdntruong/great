ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

lint:
	golangci-lint run ./... --verbose --fix

.PHONY: pre-build
pre-build:
	mkdir tmp || true
	rsync -av --progress ../commons tmp
	rsync -av --progress ../auth-ms tmp --exclude auth-ms/deployments
	cp ../go.work.sum tmp/

.PHONY: post-build
post-build:
	rm -rf tmp

.PHONY: build
build:
	docker build -t theg-auth-ms -f deployments/Dockerfile tmp/

.PHONY: bui
bui: pre-build build post-build

.PHONY: start
start:
	docker run -it --rm -p 8081:8080 --name auth-ms theg-auth-ms
