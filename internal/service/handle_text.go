package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *serv) HandleText(update tgbotapi.Update) tgbotapi.MessageConfig {

	userID := update.Message.From.ID

	// Проверяем состояние пользователя
	s.mu.Lock()
	currentState, exists := s.state[int64(userID)]
	s.mu.Unlock()

	// Обработка нажатий кнопок
	switch update.Message.Text {
	case "Создать":
		s.mu.Lock()
		s.state[int64(userID)] = "creating" // Сохраняем состояние
		s.mu.Unlock()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите текст для создания:")
		return msg

	case "Назад":
		s.mu.Lock()
		delete(s.state, int64(userID)) // Сбрасываем состояние
		s.mu.Unlock()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись назад.")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // Убираем клавиатуру
		return msg

	default:
		// Обработка следующего текстового ввода, если есть активное состояние
		if exists {
			switch currentState {
			case "creating":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Создание: "+update.Message.Text)
				s.mu.Lock()
				delete(s.state, int64(userID)) // Сбрасываем состояние
				s.mu.Unlock()
				return msg
			}
		}
		msg := s.NotFound(update)
		return msg

	}
}
