# make gen-proto src=user svc=auth
# make gen-proto src=auth svc=user
# generate protos/user.proto protobuf to services/auth-ms/internal/pkg/protos
.PHONY: gen-proto
gen-proto:
ifeq ($(svc),)
	@echo "Error: Please specify the service package using svc=<name>"
	@exit 1
endif
ifeq ($(src),)
	@echo "Error: Please specify the proto source using src=<name>"
	@exit 1
endif
	protoc --go_out=./services/$(svc)-ms/internal/pkg \
	--go-grpc_out=./services/$(svc)-ms/internal/pkg \
	--proto_path=./protos $(src).proto

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
