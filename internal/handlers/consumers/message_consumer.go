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

type MessageConsumer struct {
	telegramClient *clients.TelegramBotClient
	con            *amqp.Connection
	queueName      string
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

	tg := clients.TelegramBotClient{}

	client, _ := tg.New()

	return &MessageConsumer{con: con, queueName: env.MessagesQueueName, telegramClient: client}
}

func (consumer *MessageConsumer) Consume() {
	var forever chan struct{}
	go func() {
		ch, err := consumer.con.Channel()
		if err != nil {
			log.Fatalln("Error while create rabbitmq channel", err.Error())
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

		for message := range messages {
			// consumer.telegramClient.Send()
			m := make(map[string]string)

			err := json.Unmarshal(message.Body, &m)

			if err != nil {
				log.Fatalln("Error while trying to decode json", err)
			}

			chatId, err := strconv.ParseInt(m["chatId"], 10, 64)

			if err != nil {
				log.Fatalln("Error while trying to convert string to uint64")
			}

			consumer.telegramClient.Send(chatId, m["message"])
		}
	}()

	log.Printf("Start consuming...")
	<-forever
}
