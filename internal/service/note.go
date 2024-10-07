package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (s *serv) NoteMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	// Создаем Reply Keyboard
	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Создать"),
			tgbotapi.NewKeyboardButton("Редактировать"),
			tgbotapi.NewKeyboardButton("Удалить"),
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)

	// Отправляем сообщение с Reply Keyboard
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите опцию:")
	msg.ReplyMarkup = replyKeyboard

	// Создание inline-кнопок
	// inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "button1"),
	// 		tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "button2"),
	// 	),
	// )

	// // Отправка сообщения с inline-кнопками
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите опцию:")
	// msg.ReplyMarkup = inlineKeyboard

	return msg
}
