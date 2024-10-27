package infrastructure

import (
	"context"
	"kaisyq/tg/music/internal/infrastructure/database"
	"kaisyq/tg/music/internal/infrastructure/repositories"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Token string
}

func (telegramBot TelegramBot) Init() {

	database, error := database.GetInstance(context.Background())

	if error != nil {
		log.Fatal("error")
	}

	chatRepository := repositories.ChatRepository{Database: database}

	er := chatRepository.Insert(context.Background(), uint32(1))

	if er != nil {

	}

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

		chatRepository.Insert(context.Background(), uint32(update.Message.Chat.ID))

		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}