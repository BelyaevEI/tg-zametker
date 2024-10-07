package service

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *serv) Callback(update tgbotapi.Update) tgbotapi.MessageConfig {
	callback := update.CallbackQuery

	// Логируем нажатие для отладки
	log.Printf("CallbackQuery received: %s", callback.Data)

	// Подтверждаем callback через метод Request
	// callbackConfig := tgbotapi.NewCallback(callback.ID, "Вы выбрали: "+callback.Data)
	// if _, err := a.bot.Request(callbackConfig); err != nil {
	// 	log.Println("Error answering callback:", err)
	// }
	switch callback.Data {
	case "Создать":
		s.mu.Lock()
		s.state[int64(update.Message.From.ID)] = "creating" // Сохраняем состояние
		s.mu.Unlock()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите текст для создания заметки:")
		return msg
	case "Назад":
		s.mu.Lock()
		s.state[int64(update.Message.From.ID)] = "back" // Сохраняем состояние
		s.mu.Unlock()
		return tgbotapi.MessageConfig{}
	}
	return tgbotapi.MessageConfig{}
}
