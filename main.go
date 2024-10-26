package main

import (
	env "kaisyq/tg/music/core/env"
	"kaisyq/tg/music/infrastructure"
	"log"
)

func main() {
	var env, err = env.GetInstance()

	if err != nil {
		log.Fatalln("Error occures while trying to load environment variables", err.Error())
		panic(err.Error())
	}

	var telegramBot infrastructure.TelegramBot = infrastructure.TelegramBot{Token: env.TelegramBotToken}

	telegramBot.Init()
}
