package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	TelegramBotToken string
}

var instance *Env

func GetInstance() (*Env, error) {
	if instance != nil {
		return instance, nil
	}

	env := Env{}

	var err = env.load()

	if err != nil {
		return nil, err
	}

	return &env, nil
}

func (env *Env) load() error {

	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Cannot load .env file")
	}

	var telegramBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")

	if telegramBotToken == "" {
		log.Fatalln("Cannot find value by TELEGRAM_BOT_TOKEN key")

		return fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	env.TelegramBotToken = telegramBotToken

	return nil
}
