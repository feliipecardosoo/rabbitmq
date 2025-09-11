# =========================
# Variáveis
# =========================
PROJECT_NAME   = rabbitmq-ms
APP_IMAGE_NAME = feliipecardosoo/go-app-ms
APP_VERSION    = latest
FULL_APP_IMAGE = $(APP_IMAGE_NAME):$(APP_VERSION)

# =========================
# Push da imagem Docker
# =========================
push:
	docker push $(FULL_APP_IMAGE)

# =========================
# Build Docker
# =========================
build:
	docker build -t $(FULL_APP_IMAGE) .

docker-compose-up:
	docker compose -p $(PROJECT_NAME) up -d

docker-compose-down:
	docker compose -p $(PROJECT_NAME) down

# =========================
# Sobe o RabbitMQ + app
# =========================
up: build docker-compose-down docker-compose-up

down:
	docker compose -p $(PROJECT_NAME) down

restart: down up

# =========================
# Logs
# =========================
logs:
	docker compose -p $(PROJECT_NAME) logs -f

ps:
	docker compose -p $(PROJECT_NAME) ps

# =========================
# RabbitMQ healthcheck
# =========================
rabbitmq:
	docker exec -it rabbitmq rabbitmq-diagnostics ping

# =========================
# Remove containers, imagens e volumes
# =========================
clean:
	docker compose -p $(PROJECT_NAME) down -v --rmi all --remove-orphans
