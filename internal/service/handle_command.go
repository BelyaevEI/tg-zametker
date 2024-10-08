package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (s *serv) handleCommnds(update tgbotapi.Update, state string) tgbotapi.MessageConfig {

	switch state {
	case "creating":

		err := s.repository.CreateNote(update.Message.From.ID, update.Message.Text)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "При создании возникла ошибка, попробуйте снова")
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Создание заметки успешно")
		s.mu.Lock()
		delete(s.state, int64(update.Message.From.ID)) // Сбрасываем состояние
		s.mu.Unlock()
		return msg
	}

	return tgbotapi.MessageConfig{}
}
