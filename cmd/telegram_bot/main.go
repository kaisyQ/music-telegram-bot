package main

import (
	"fmt"
	"kaisyq/tg/music/internal/core/env"
	"kaisyq/tg/music/internal/handlers"
	message_consumer "kaisyq/tg/music/internal/handlers/consumers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	env, err := env.GetInstance()

	if err != nil {
		log.Println("Cannot load .env file")
	}

	telegramBotHandler := handlers.New()

	router := gin.Default()

	{
		v1 := router.Group("/v1")

		telegram := v1.Group("/telegram-bot")

		telegram.POST("/handle", telegramBotHandler.Handle)

	}

	messageConsumer := message_consumer.New()

	go messageConsumer.Consume()

	router.Run(fmt.Sprintf(":%d", env.RestApiPort))

	messageConsumer.Close()
}
