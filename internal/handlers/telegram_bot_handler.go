package handlers

import (
	"kaisyq/tg/music/internal/infrastructure/clients"
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotHandler struct {
	telegramClient *clients.TelegramBotClient
}

var instance *TelegramBotHandler

func New() *TelegramBotHandler {

	tg := clients.TelegramBotClient{}

	client, _ := tg.New()

	if instance == nil {
		instance = &TelegramBotHandler{telegramClient: client}
	}

	return instance
}

func (handler TelegramBotHandler) Handle(c *gin.Context) {

	var update tgbotapi.Update

	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.telegramClient.Send(uint64(update.Message.Chat.ID), "Новая обработка сообщения!")

}
