ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: lint
lint:
	golangci-lint run ./... --verbose

.PHONY: build
build:
	go build ./cmd/api

# Docker usage | https://github.com/golang-migrate/migrate?tab=readme-ov-file#docker-usage
.PHONY: db-migrate
db-migrate:
	docker run --rm -v $(ROOT_DIR)/db/migrations:/migrations --network host migrate/migrate \
	-path=/migrations -database $(LOCAL_DATABASE_URL) up

.PHONY: db-migrate-up
db-migrate-up:
	docker run --rm -v $(ROOT_DIR)/db/migrations:/migrations --network host migrate/migrate \
    -path=/migrations -database $(LOCAL_DATABASE_URL) up 1

.PHONY: db-migrate-down
db-migrate-down:
	docker run --rm -v $(ROOT_DIR)/db/migrations:/migrations --network host migrate/migrate \
	-path=/migrations -database $(LOCAL_DATABASE_URL) down 1
