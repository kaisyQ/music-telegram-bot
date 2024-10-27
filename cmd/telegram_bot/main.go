package main

import (
	env "kaisyq/tg/music/internal/core/env"
	"kaisyq/tg/music/internal/infrastructure"
	"log"
)

func main() {
	var env, err = env.GetInstance()

	if err != nil {
		log.Fatalln("Error occures while trying to load environment variables", err.Error())
		panic(err.Error())
	}

	// router := gin/

	var telegramBot infrastructure.TelegramBot = infrastructure.TelegramBot{Token: env.TelegramBotToken}

	telegramBot.Init()
}
