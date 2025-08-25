package main

import (
	"log"
	"rabbitmq/src/config/env"
	"rabbitmq/src/config/rabbitmq"
)

func main() {
	// Carrega as variáveis de ambiente do projeto (.env)
	env.LoadEnv()

	// --------------------------
	// Conexão com RabbitMQ
	// --------------------------
	rm, err := rabbitmq.NewRabbitMQConnection()
	if err != nil {
		log.Fatal("Erro ao conectar ao RabbitMQ:", err)
	}
	// Garante que a conexão será fechada ao finalizar o programa
	defer func() {
		if err := rm.Close(); err != nil {
			log.Println("Erro ao fechar RabbitMQ:", err)
		}
	}()

	// --------------------------
	// Publicação de mensagens
	// --------------------------

	// Publica mensagem correta na fila pré-definida
	if err := rm.Publish(rabbitmq.QueueEmails, "Bem-vindo!"); err != nil {
		log.Println("Erro ao publicar mensagem na fila:", err)
	} else {
		log.Println("Mensagem publicada com sucesso na fila:", rabbitmq.QueueEmails)
	}

	// Tenta publicar em uma fila que não está permitida (vai gerar erro)
	if err := rm.Publish("fila_nao_existente", "Teste"); err != nil {
		log.Println("Erro esperado ao publicar em fila inválida:", err)
	}

}
