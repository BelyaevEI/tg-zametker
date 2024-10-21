package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *serv) HandleText(update tgbotapi.Update) tgbotapi.MessageConfig {

	// Проверяем состояние пользователя
	s.mu.Lock()
	currentState, exists := s.state[int64(update.Message.From.ID)]
	s.mu.Unlock()

	// Обработка нажатий кнопок
	switch update.Message.Text {
	case "Создать":
		s.mu.Lock()
		s.state[int64(update.Message.From.ID)] = "creating" // Сохраняем состояние
		s.mu.Unlock()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите текст для создания:")

		return msg
	case "Показать заметки":
		list, err := s.showNotes(update.Message.From.ID)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Возникла ошибка, попробуйте снова.")
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, list)

		return msg
	case "Удалить":
		s.mu.Lock()
		s.state[int64(update.Message.From.ID)] = "delete" // Сохраняем состояние
		s.mu.Unlock()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите номер заметки для удаления")

		return msg
	case "Назад":
		s.mu.Lock()
		delete(s.state, int64(update.Message.From.ID)) // Сбрасываем состояние
		s.mu.Unlock()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы вернулись назад.")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // Убираем клавиатуру

		return msg
	default:

		// Обработка следующего текстового ввода, если есть активное состояние
		if exists {
			msg := s.handleCommnds(update, currentState)
			return msg
		}
		// если вводится в чат текст без команды
		msg := s.NotFound(update)

		return msg
	}
}
