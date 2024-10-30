package main

import (
	"kaisyq/tg/music/internal/handlers"
	message_consumer "kaisyq/tg/music/internal/handlers/consumers"

	"github.com/gin-gonic/gin"
)

func main() {
	telegramBotHandler := handlers.New()

	router := gin.Default()

	{
		v1 := router.Group("/v1")

		telegram := v1.Group("/telegram-bot")

		telegram.POST("/handle", telegramBotHandler.Handle)

	}

	go func() {
		message_consumer.New().Consume()
	}()

	router.Run(":8000")

}
