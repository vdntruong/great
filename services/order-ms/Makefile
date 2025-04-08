ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
TEMPORARY_DIR := tmp
CONTAINER_NAME := order-ms
DOCKER_IMAGE_NAME := theg-order-ms

.PHONY: lint
lint:
	golangci-lint run ./... --verbose

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

.PHONY: db-seed
db-seed:
	docker run --rm -v $(ROOT_DIR)/db/seeds:/seeds --network host migrate/migrate \
	-path=/seeds -database $(LOCAL_DATABASE_URL) up

.PHONY: db-seed-down
db-seed-down:
	docker run --rm -v $(ROOT_DIR)/db/seeds:/seeds --network host migrate/migrate \
	-path=/seeds -database $(LOCAL_DATABASE_URL) down

## Build docker image

.PHONY: setup-tmp
setup-tmp:
	mkdir $(TEMPORARY_DIR) || true
	rsync -av --progress ../commons $(TEMPORARY_DIR)
	rsync -av --progress ../order-ms $(TEMPORARY_DIR) --exclude order-ms/deployments
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
	docker run -it --rm -p 8084:8080 --name $(CONTAINER_NAME) $(DOCKER_IMAGE_NAME)

.PHONY: create-migration
create-migration:
	migrate create -ext sql -dir db/migrations -seq $(NEW_FILE)
