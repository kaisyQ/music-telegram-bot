package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	TelegramBotToken string
	DatabaseUrl      string
	RabbitmqUser     string
	RabbitmqPassword string
	RabbitmqHost     string
	RabbitmqPort     string

	MessagesQueueName string

	RestApiPort int
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
		log.Fatalln("Cannot load .env file", err.Error())
	}

	restApiPortString := os.Getenv("REST_API_PORT")

	if restApiPortString == "" {
		return printAndReturnError("REST_API_PORT")
	}

	restApiPort, err := strconv.Atoi(restApiPortString)

	if err != nil {
		log.Fatalln("Error while trying to parse rest api port to int32", err.Error())
	}

	env.RestApiPort = restApiPort

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	if telegramBotToken == "" {
		return printAndReturnError("TELEGRAM_BOT_TOKEN")
	}

	env.TelegramBotToken = telegramBotToken

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		return printAndReturnError("DATABASE_URL")
	}

	env.DatabaseUrl = databaseUrl

	rabbitmqUser := os.Getenv("RABBIT_MQ_USER")

	if rabbitmqUser == "" {
		return printAndReturnError("RABBIT_MQ_USER")
	}

	env.RabbitmqUser = rabbitmqUser

	rabbitmqPassword := os.Getenv("RABBIT_MQ_PASSWORD")

	if rabbitmqPassword == "" {
		return printAndReturnError("RABBIT_MQ_PASSWORD")
	}

	env.RabbitmqPassword = rabbitmqPassword

	rabbitmqHost := os.Getenv("RABBIT_MQ_HOST")

	if rabbitmqHost == "" {
		return printAndReturnError("RABBIT_MQ_HOST")
	}

	env.RabbitmqHost = rabbitmqHost

	rabbitmqPort := os.Getenv("RABBIT_MQ_PORT")

	if rabbitmqPort == "" {
		return printAndReturnError("RABBIT_MQ_PORT")
	}

	env.RabbitmqPort = rabbitmqPort

	messagesQueueName := os.Getenv("MESSAGES_QUEUE_NAME")

	if messagesQueueName == "" {
		return printAndReturnError("MESSAGES_QUEUE_NAME")
	}

	env.MessagesQueueName = messagesQueueName

	return nil
}

func printAndReturnError(key string) error {
	log.Fatalf("Cannot find value by %s key", key)

	return fmt.Errorf("%s environment variable is not set", key)
}
