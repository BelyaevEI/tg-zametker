package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (s *serv) HandleText(update tgbotapi.Update) tgbotapi.MessageConfig {

	// Проверяем состояние пользователя
	s.mu.Lock()
	currentState, exists := s.state[int64(update.Message.From.ID)]
	s.mu.Unlock()

	//если есть состояние
	if exists {
		switch currentState {
		case "creating":
		}
	}

	msg := s.NotFound(update)
	return msg

}
