ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

lint:
	golangci-lint run ./... --verbose --fix

.PHONY: pre-build
pre-build:
	mkdir tmp || true
	rsync -av --progress ../commons tmp
	rsync -av --progress ../user-ms tmp --exclude user-ms/deployments
	cp ../go.work.sum tmp/

.PHONY: post-build
post-build:
	rm -rf tmp

.PHONY: build
build:
	docker build -t theg-user-ms -f deployments/Dockerfile tmp/

.PHONY: bui
bui: pre-build build post-build

.PHONY: start
start:
	docker run -it --rm -p 8082:8080 --name user-ms theg-user-ms

.PHONY: live
live:
	air


# Docker usage | https://github.com/golang-migrate/migrate?tab=readme-ov-file#docker-usage
.PHONY: db-local-migrate
db-local-migrate:
	docker run --rm -v $(ROOT_DIR)/db/migrations:/migrations --network host migrate/migrate \
	-path=/migrations -database $(LOCAL_DATABASE_URL) up
