ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

TEMPORARY_DIR := tmp
CONTAINER := user-ms
DOCKER_IMAGE := theg-user-ms

.PHONY: lint
lint:
	golangci-lint run ./... --verbose

.PHONY: setup-tmp
setup-tmp:
	mkdir $(TEMPORARY_DIR) || true
	rsync -av --progress ../commons $(TEMPORARY_DIR)
	rsync -av --progress ../user-ms $(TEMPORARY_DIR) --exclude user-ms/deployments
	cp ../go.work.sum $(TEMPORARY_DIR)/

.PHONY: clean-tmp
clean-tmp:
	rm -rf $(TEMPORARY_DIR)

.PHONY: docker-image
docker-image:
	docker build -t $(DOCKER_IMAGE) -f deployments/Dockerfile $(TEMPORARY_DIR)/

.PHONY: build
build: clean-tmp setup-tmp docker-image clean-tmp

.PHONY: start-container
start-container:
	docker run -it --rm -p 8082:8080 --name $(CONTAINER) $(DOCKER_IMAGE)

.PHONY: start-live
start-live:
	air

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
