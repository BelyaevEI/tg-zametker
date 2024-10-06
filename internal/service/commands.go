package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обработка всех команд для бота
func (s *serv) Commands(update tgbotapi.Update) tgbotapi.MessageConfig {
	switch update.Message.Command() {
	case "note":
		return s.NoteMenu(update)
	case "notification":
		return tgbotapi.MessageConfig{}
	case "info":
		return tgbotapi.MessageConfig{}
	case "start":
		return s.Start(update)
	default:
		return s.NotFound(update)
	}
}
