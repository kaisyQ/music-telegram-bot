package message_consumer

import (
	"kaisyq/tg/music/internal/config"
	"kaisyq/tg/music/internal/core/env"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageConsumer struct {
	con       *amqp.Connection
	queueName string
}

func New() *MessageConsumer {
	env, err := env.GetInstance()

	if err != nil {
		log.Fatalln("Cannot load .env file")
	}

	rabbitUrl := config.GetRabbitUrl()

	con, err := amqp.Dial(rabbitUrl)

	if err != nil {
		log.Fatalln("Error while creating connect to queue")
	}
	return &MessageConsumer{con: con, queueName: env.MessagesQueueName}
}

func (consumer *MessageConsumer) Consume() {
	ch, err := consumer.con.Channel()

	if err != nil {
		log.Fatalln("Error while create rabbitmq channel", err)
	}

	//@TODO need to fix this closing connections(try to found try-with-resources interface alternative)

	defer consumer.con.Close()
	defer ch.Close()

	messages, err := ch.Consume(
		consumer.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("failed to register a consumer. Error: %s", err)
	}

	var forever chan struct{}
	go func() {
		for message := range messages {
			log.Printf("received a message: %s", message.Body)
		}
	}()

	log.Printf("Start consuming...")
	<-forever
}
