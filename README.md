# Exemplo RabbitMQ

Este repositório é uma amostra simples de uso do RabbitMQ, criado para meu primeiro contato com a ferramenta.

## Como usar

1. Certifique-se de ter o [Docker](https://www.docker.com/) e o [Docker Compose](https://docs.docker.com/compose/) instalados.
2. Execute o comando abaixo na pasta do projeto para iniciar o RabbitMQ:

   ```sh
   docker-compose up -d
    ```

3. Acesse a interface de gerenciamento do RabbitMQ em [http://localhost:15673](http://localhost:15673) com as credenciais:
   - Usuário: `admin`
   - Senha: `admin`

As portas padrão foram alteradas para evitar conflitos:
- Porta do RabbitMQ: `5673` (em vez de `5672`)
- Porta da interface de gerenciamento: `15673` (em vez de `15672`)

## Observações
- Esta configuração é para fins de aprendizado e desenvolvimento. Para ambientes de produção, considere ajustar as configurações de segurança e persistência de dados.
