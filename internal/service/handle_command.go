package service

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *serv) handleCommnds(update tgbotapi.Update, state string) tgbotapi.MessageConfig {

	switch state {
	case "creating":
		err := s.repository.CreateNote(update.Message.From.ID, update.Message.Text)
		if err != nil {
			log.Printf("creating is failed: %v", err)
			return tgbotapi.NewMessage(update.Message.Chat.ID, "При создании возникла ошибка, попробуйте снова.")
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Создание заметки успешно.")
		s.mu.Lock()
		delete(s.state, int64(update.Message.From.ID)) // Сбрасываем состояние
		s.mu.Unlock()

		return msg
	case "delete":
		err := s.repository.DeleteNote(update.Message.From.ID, update.Message.Text)
		if err != nil {
			log.Printf("deleting is failed: %v", err)
			return tgbotapi.NewMessage(update.Message.Chat.ID, "При удалении возникла ошибка, попробуйте снова.")
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Удаление заметки успешно.")
		s.mu.Lock()
		delete(s.state, int64(update.Message.From.ID)) // Сбрасываем состояние
		s.mu.Unlock()

		return msg
	case "input":
		numNote, err := strconv.Atoi(update.Message.Text)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Введите правльный номер заметки.")
		}

		s.numberNote = int64(numNote) // Сохраним номер редактируемой заметки

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите текст новой заметки:")
		s.mu.Lock()
		delete(s.state, int64(update.Message.From.ID))  // Сбрасываем состояние
		s.state[int64(update.Message.From.ID)] = "edit" // Сохраняем состояние
		s.mu.Unlock()

		return msg
	case "edit":
		err := s.repository.EditNote(update.Message.From.ID, s.numberNote, update.Message.Text)
		if err != nil {
			log.Printf("updating is failed: %v", err)
			return tgbotapi.NewMessage(update.Message.Chat.ID, "При редактировании возникла ошибка, попробуйте снова.")
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Редактирование заметки успешно.")
		s.mu.Lock()
		delete(s.state, int64(update.Message.From.ID)) // Сбрасываем состояние
		s.mu.Unlock()

		return msg
	}

	return tgbotapi.MessageConfig{}
}
