ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
PROJECT_NAME := theg

# docker compose
COMPOSE_FILE := ./deployments/docker/docker-compose.yml

# services builds

.PHONY: build-auth
build-auth:
	make -C $(ROOT_DIR)/services/auth-ms build

.PHONY: build-user
build-user:
	make -C $(ROOT_DIR)/services/user-ms build

.PHONY: build-product
build-product:
	make -C $(ROOT_DIR)/services/product-ms build

.PHONY: build-services
build-services:
	@make build-auth & make build-user & make build-product & wait

# build service images and start docker compose

.PHONY: build-up
build-up: build-services
	docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d

# just start docker compose

.PHONY: up
up:
	docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d

.PHONY: down
down:
	docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down

.PHONY: logs
logs:
	docker compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) logs -f
