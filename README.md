# Exemplo RabbitMQ em Go

Este repositório é uma amostra simples de uso do RabbitMQ com Go, criado para meu primeiro contato com a ferramenta e integração via Docker.

## Sobre

O projeto demonstra como publicar mensagens em filas do RabbitMQ utilizando uma aplicação escrita em Go. O ambiente é totalmente orquestrado com Docker Compose, facilitando a execução local sem necessidade de instalações manuais.

## Como executar

1. **Pré-requisitos:**  
   Tenha [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) instalados.

2. **Suba os containers:**  
   Na raiz do projeto, execute:

   ```sh
   docker-compose up --build -d
   ```

3 . **Acesse o RabbitMQ:**
Interface de gerenciamento: [http://localhost:15673]( http://localhost:15673 )

   Usuário: `admin`  
   Senha: `admin`

   **Portas expostas:**
5673: Porta principal do RabbitMQ (amqp)
15673: Interface de gerenciamento

4 . **Logs da aplicação Go:**
Para acompanhar os logs da aplicação:

```sh
docker-compose logs -f app
```

## Estrutura do Projeto

```

rabbitmq/
├── Dockerfile
├── docker-compose.yaml
├── go.mod / go.sum
├── main.go
├── src/
│   └── config/
│       ├── env/
│       │   └── env.go
│       └── rabbitmq/
│           └── connection.go
└── [README.md](http://_vscodecontentref_/0)
```

## Observações

- Este projeto é apenas para fins educacionais e de aprendizado.
- Sinta-se à vontade para expandir e modificar conforme necessário.
