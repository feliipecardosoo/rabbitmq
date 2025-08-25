package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ encapsula a conexão e o canal com o servidor RabbitMQ.
type RabbitMQ struct {
	Conn    *amqp.Connection // Conexão com o RabbitMQ
	Channel *amqp.Channel    // Canal usado para publicar e consumir mensagens
}

// Nomes de filas pré-definidos para evitar o uso de strings arbitrárias.
// Pode ser configurado via variável de ambiente, caso não exista usa valor padrão.
// Isso garante padronização e segurança ao trabalhar com filas.
var (
	QueueEmails  = "emails_queue"  // Fila para envio de e-mails
	QueueMembers = "members_queue" // Fila para processamento de membros
)

// allowedQueues lista todas as filas válidas do sistema.
// A validação é feita antes de publicar ou consumir mensagens.
var allowedQueues = map[string]bool{
	QueueEmails:  true,
	QueueMembers: true,
}

// NewRabbitMQConnection cria e retorna uma nova conexão com o RabbitMQ.
// Lê a URI de conexão da variável de ambiente RABBITMQ_URI.
// Retorna um ponteiro para RabbitMQ e erro, caso haja falha.
func NewRabbitMQConnection() (*RabbitMQ, error) {
	uri := os.Getenv("RABBITMQ_URI")
	if uri == "" {
		log.Fatal("RABBITMQ_URI não configurado")
	}

	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	return &RabbitMQ{Conn: conn, Channel: ch}, nil
}

// Close encerra o canal e a conexão com o RabbitMQ.
// Deve ser chamado ao final do uso para liberar recursos.
func (r *RabbitMQ) Close() error {
	if err := r.Channel.Close(); err != nil {
		return fmt.Errorf("failed to close channel: %w", err)
	}
	if err := r.Conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}

// validateQueue verifica se a fila fornecida está na lista de filas permitidas.
// Retorna erro caso a fila não seja válida.
func validateQueue(queueName string) error {
	if !allowedQueues[queueName] {
		return errors.New("fila não permitida: " + queueName)
	}
	return nil
}

// Publish envia uma mensagem para uma fila especificada.
// - Valida se a fila é permitida antes de publicar.
// - Cria a fila caso ela ainda não exista.
// - Usa um contexto com timeout de 5 segundos para evitar bloqueios.
func (r *RabbitMQ) Publish(queueName, message string) error {
	if err := validateQueue(queueName); err != nil {
		return err
	}

	_, err := r.Channel.QueueDeclare(
		queueName, // nome da fila
		true,      // durable: persiste após reinício do broker
		false,     // auto-delete: não apaga automaticamente
		false,     // exclusive: não exclusivo para a conexão
		false,     // no-wait: aguarda confirmação
		nil,       // argumentos adicionais
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Channel.PublishWithContext(
		ctx,
		"",        // exchange (default)
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

// Consume inicia o consumo de mensagens de uma fila especificada.
// - Valida se a fila é permitida antes de consumir.
// - Cria a fila caso ela ainda não exista.
// Retorna um canal de mensagens (amqp.Delivery) para ser percorrido.
func (r *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	if err := validateQueue(queueName); err != nil {
		return nil, err
	}

	_, err := r.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	msgs, err := r.Channel.Consume(
		queueName, // nome da fila
		"",        // consumer tag (gerado automaticamente)
		true,      // auto-ack: confirma recebimento automaticamente
		false,     // exclusive: não exclusivo
		false,     // no-local (ignorado pelo RabbitMQ)
		false,     // no-wait: aguarda confirmação
		nil,       // argumentos adicionais
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume: %w", err)
	}

	return msgs, nil
}
