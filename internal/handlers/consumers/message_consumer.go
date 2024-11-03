package message_consumer

import (
	"encoding/json"
	"kaisyq/tg/music/internal/config"
	"kaisyq/tg/music/internal/core/env"
	"kaisyq/tg/music/internal/infrastructure/clients"
	"log"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

// @TODO Create more flexible structure for working with consumers. Now its look like shit
type MessageConsumer struct {
	telegramClient *clients.TelegramBotClient
	con            *amqp.Connection
	queueName      string
}

func New() *MessageConsumer {
	env, err := env.GetInstance()

	if err != nil {
		log.Fatalln("Cannot load .env file", err.Error())
	}

	rabbitUrl := config.GetRabbitUrl()

	con, err := amqp.Dial(rabbitUrl)

	if err != nil {
		log.Fatalln("Error while creating connect to queue", err.Error())
	}

	tg := clients.TelegramBotClient{}

	client, _ := tg.New()

	return &MessageConsumer{con: con, queueName: env.MessagesQueueName, telegramClient: client}
}

func (consumer *MessageConsumer) Consume() {
	ch, err := consumer.con.Channel()
	var forever chan struct{}
	go func() {
		if err != nil {
			log.Fatalln("Error while create rabbitmq channel", err.Error())
		}

		messages, err := ch.Consume(
			consumer.queueName,
			"",
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			log.Fatalf("failed to register a consumer. Error: %s", err.Error())
		}

		for message := range messages {
			m := make(map[string]string)

			err := json.Unmarshal(message.Body, &m)

			if err != nil {
				log.Fatalln("Error while trying to decode json", err.Error())
			}

			chatId, err := strconv.ParseInt(m["chatId"], 10, 64)

			if err != nil {
				log.Fatalln("Error while trying to convert string to uint64", err.Error())
			}

			consumer.telegramClient.Send(chatId, m["message"])
		}
	}()

	log.Printf("Start consuming %s...", consumer.queueName)
	<-forever
}

func (consumer *MessageConsumer) Close() {
	defer consumer.con.Close()
}
