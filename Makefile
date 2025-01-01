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
