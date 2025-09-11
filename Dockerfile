# =========================
# Build da aplicação
# =========================
FROM golang:1.24.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app main.go

# =========================
# Imagem final
# =========================
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

# Variável de ambiente para conexão com RabbitMQ
ENV RABBITMQ_URI=amqp://admin:admin@rabbitmq:5672/

CMD ["./app"]
