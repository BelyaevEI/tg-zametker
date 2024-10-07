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
			log.Printf("Update received: %+v", update) // Лог всех обновлений

			if update.CallbackQuery != nil {
				msg := a.serviceProvider.service.Callback(update)
				a.bot.Send(msg)
			}

			if update.Message == nil {
				continue
			}

			// если пришла команда
			if update.Message.IsCommand() {
				msg := a.serviceProvider.service.Commands(update)
				a.bot.Send(msg)
			} else {
				msg := a.serviceProvider.service.HandleText(update)
				a.bot.Send(msg)
			}

		case <-ctx.Done():
			log.Println("Shutting down the update listener")
			return nil
		}
	}
}
