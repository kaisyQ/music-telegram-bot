package infrastructure

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Token string
}

func (telegramBot TelegramBot) Init() {
	bot, err := tgbotapi.NewBotAPI(telegramBot.Token)

	if err != nil {
		log.Fatal("Cannot initialize telegram bot from token")
	}

	bot.Debug = true

	log.Printf("Running telegram bot with name %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
