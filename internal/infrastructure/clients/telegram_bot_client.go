package clients

import (
	"kaisyq/tg/music/internal/core/env"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// @TODO make thread safe
type TelegramBotClient struct {
	api *tgbotapi.BotAPI
}

var instance *TelegramBotClient

func (client TelegramBotClient) New() (*TelegramBotClient, error) {

	env, err := env.GetInstance()

	if err != nil {
		log.Fatalln("Error while trying to get .env")

		return nil, err
	}

	api, err := tgbotapi.NewBotAPI(env.TelegramBotToken)

	if err != nil {
		log.Fatalln("Error while trying to create bot telegram api bot instance")

		return nil, err
	}

	if instance == nil {
		instance = &TelegramBotClient{api: api}
	}

	return instance, nil
}

func (client TelegramBotClient) Send(chatId uint64, messageText string) error {

	message := tgbotapi.NewMessage(int64(chatId), messageText)

	_, err := client.api.Send(message)

	if err != nil {
		return err
	}

	return nil
}
