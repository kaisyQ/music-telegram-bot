package config

import (
	"fmt"
	"kaisyq/tg/music/internal/core/env"
	"log"
)

func GetRabbitUrl() string {
	env, err := env.GetInstance()

	if err != nil {
		log.Fatalln("Cannot load .env file")
	}

	return fmt.Sprintf(
		getRabbitConTemplate(),
		env.RabbitmqUser,
		env.RabbitmqPassword,
		env.RabbitmqHost,
		env.RabbitmqPort,
	)
}

func getRabbitConTemplate() string {
	return "amqp://%s:%s@%s:%s"
}
