package handlers

import (
	"encoding/json"
	message_producer "kaisyq/tg/music/internal/infrastructure/producers"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBotHandler struct {
	messageProducer *message_producer.MessageProducer
}

var instance *TelegramBotHandler

func New() *TelegramBotHandler {
	messageProducer := message_producer.New()

	if instance == nil {
		instance = &TelegramBotHandler{
			messageProducer: messageProducer,
		}
	}

	return instance
}

func (handler TelegramBotHandler) Handle(c *gin.Context) {

	var update tgbotapi.Update

	if err := c.BindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := make(map[string]string)

	m["Message"] = "Новая обработка сообщний через очередь"
	m["chatId"] = strconv.Itoa(int(update.Message.Chat.ID))

	message, err := json.Marshal(m)

	if err != nil {
		log.Fatalln("Error while trying to encode json from map", err)
	}

	handler.messageProducer.Produce(message)
}
