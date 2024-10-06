package app

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (a *App) UpdatesFromTGServer(ctx context.Context) error {

	// Устанавливаем команды для меню бота
	commands := []tgbotapi.BotCommand{
		{
			Command:     "note",
			Description: "Заметка",
		},
		{
			Command:     "notification",
			Description: "Уведомление",
		},
		{
			Command:     "info",
			Description: "Информация о боте",
		},
	}

	// Установка команд
	if _, err := a.bot.Request(tgbotapi.NewSetMyCommands(commands...)); err != nil {
		return err
	}

	// Получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := a.bot.GetUpdatesChan(u)
	for {
		select {
		case update := <-updates:
			// проверяем, что сообщение не пустое
			if update.Message == nil {
				continue
			}

			// если пришла команда
			if update.Message.IsCommand() {
				msg := a.serviceProvider.service.Commands(update)
				a.bot.Send(msg)
			} else {
				// Обработка нажатий кнопок
				switch update.Message.Text {
				case "Кнопка 1":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали кнопку 1")
					a.bot.Send(msg)

				case "Кнопка 2":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали кнопку 2")
					a.bot.Send(msg)
				default:
					msg := a.serviceProvider.service.NotFound(update)
					a.bot.Send(msg)
				}
			}

		case <-ctx.Done():
			log.Println("Shutting down the update listener")
			return nil
		}
	}
}
