package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *serv) Start(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("Привет %v. Для взаимодействия с ботом выберите команду из меню.",
			update.Message.Chat.FirstName))
}
