package message_producer

import (
	"context"
	"kaisyq/tg/music/internal/config"
	"kaisyq/tg/music/internal/core/env"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageProducer struct {
	con       *amqp.Connection
	queueName string
}

func New() *MessageProducer {
	env, err := env.GetInstance()

	if err != nil {
		log.Fatalln("Cannot load .env file")
	}

	rabbitUrl := config.GetRabbitUrl()

	con, err := amqp.Dial(rabbitUrl)

	if err != nil {
		log.Fatalln("Error while creating connect to queue")
	}

	return &MessageProducer{con: con, queueName: env.MessagesQueueName}
}

func (producer *MessageProducer) Produce(message string) {
	defer producer.con.Close()

	ch, err := producer.con.Channel()

	if err != nil {
		log.Fatalln("Error while create rabbitmq channel")
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		producer.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue. Error: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err)
	}

	log.Printf(" [x] Sent %s\n", message)
}
