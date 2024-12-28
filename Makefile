# golang service

proto-auth:
	protoc --proto_path=./protos \
	--go-grpc_out=./services/auth-ms/pkg \
	--micro_out=./services/auth-ms/pkg \
	--go_out=./services/auth-ms/pkg auth.proto

proto-user:
	protoc --proto_path=./protos \
	--go-grpc_out=./services/user-ms/internal/pkg \
	--go_out=./services/user-ms/internal/pkg user.proto

# docker compose

COMPOSE_FILE := ./deployments/docker/docker-compose.yml
COMPOSE_PROJECT_NAME := theg

.PHONY: up
up:
	docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROJECT_NAME) up -d

.PHONY: down
down:
	docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROJECT_NAME) down

.PHONY: logs
logs:
	docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROJECT_NAME) logs -f
