# App name
APP_NAME = server

# Goose setting
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING = "admin:mysql@tcp(127.0.0.1:8811)/go-ec"
GOOSE_MIGRATION_DIR ?= sql/schemas

dev:
	docker start qdvn-redis && docker start qdvn-mysql-master && CONFIG_NAME=local go run ./cmd/${APP_NAME}/

dev_down:
	docker stop qdvn-redis && docker stop qdvn-mysql-master

run:
	docker start qdvn-redis && docker start qdvn-mysql-master && docker compose up -d && go run ./cmd/$(APP_NAME)

kill:
	docker compose kill

up:
	docker start qdvn-redis && docker start qdvn-mysql-master \
	&& until docker exec qdvn-mysql-master mysqladmin ping -h "localhost" --silent; do echo "Waiting for MySQL..."; sleep 2; done \
	&& until docker exec qdvn-redis redis-cli ping | grep PONG > /dev/null; do echo "Waiting for Redis..."; sleep 2; done \
	&& docker compose up -d

down:
	docker compose down && docker stop qdvn-redis && docker stop qdvn-mysql-master

.PHONY: run

.PHONE: air

goose_up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

goose_down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

goose_reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset