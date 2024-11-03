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

func (producer *MessageProducer) Produce(message []byte) {
	ch, err := producer.con.Channel()

	if err != nil {
		log.Fatalln("Error while create rabbitmq channel", err.Error())
	}

	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",
		producer.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)

	if err != nil {
		log.Fatalf("failed to publish a message. Error: %s", err.Error())
	}

	log.Printf("[x] Sent %s\n", message)
}

func (producer *MessageProducer) Close() {
	defer producer.con.Close()
}
