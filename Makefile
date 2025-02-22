# App name
APP_NAME = server

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