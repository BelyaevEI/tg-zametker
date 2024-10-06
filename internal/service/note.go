package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (s *serv) NoteMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	// Создаем Reply Keyboard
	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Кнопка 1"),
			tgbotapi.NewKeyboardButton("Кнопка 2"),
		),
	)

	// Отправляем сообщение с Reply Keyboard
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите опцию:")
	msg.ReplyMarkup = replyKeyboard

	return msg
}
