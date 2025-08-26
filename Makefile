PROJECT_NAME = sistema-ibt
IMAGE_NAME = rabbitmq-sistema-ibt
NETWORK_NAME = sistema-ibt-network
RABBIT_CONTAINER = rabbitmq-sistema-ibt

# Cria a rede se não existir
network:
	@docker network inspect $(NETWORK_NAME) >/dev/null 2>&1 || \
	docker network create $(NETWORK_NAME)

# Builda a imagem customizada
build: network
	docker build -t $(IMAGE_NAME) .

# Sobe os containers via docker-compose
up: build
	docker compose -p $(PROJECT_NAME) up -d 
	
# Derruba os containers
down:
	docker compose -p $(PROJECT_NAME) down

# Restart dos containers
restart: down up

# Mostra os logs
logs:
	docker compose -p $(PROJECT_NAME) logs -f

# Lista containers ativos
ps:
	docker compose -p $(PROJECT_NAME) ps

# Remove containers, imagens e volumes do projeto
clean:
	docker compose -p $(PROJECT_NAME) down -v --rmi all --remove-orphans
	@docker network rm $(NETWORK_NAME) 2>/dev/null || true

# Entrar no container RabbitMQ
shell-rabbit:
	docker exec -it $(RABBIT_CONTAINER) bash
