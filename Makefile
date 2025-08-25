PROJECT_NAME = rabbitmq-ms

# Atalhos para docker compose
up:
	docker compose -p $(PROJECT_NAME) up -d

down:
	docker compose -p $(PROJECT_NAME) down

build:
	docker compose -p $(PROJECT_NAME) build

restart: down up

logs:
	docker compose -p $(PROJECT_NAME) logs -f

ps:
	docker compose -p $(PROJECT_NAME) ps

rabbitmq:
	docker exec -it rabbitmq rabbitmq-diagnostics ping

clean:
	docker compose -p $(PROJECT_NAME) down -v --rmi all --remove-orphans
