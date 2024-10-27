package main

import (
	"kaisyq/tg/music/internal/handlers"

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

	router.Run(":8000")
}
