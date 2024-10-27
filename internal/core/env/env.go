package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	TelegramBotToken string
	DatabaseUrl      string
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
		return printAndReturnError("TELEGRAM_BOT_TOKEN")
	}

	env.TelegramBotToken = telegramBotToken

	var databaseUrl = os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		return printAndReturnError("DATABASE_URL")
	}

	env.DatabaseUrl = databaseUrl

	return nil
}

func printAndReturnError(key string) error {
	log.Fatalf("Cannot find value by %s key", key)

	return fmt.Errorf("%s environment variable is not set", key)
}
