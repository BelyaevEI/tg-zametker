package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Получаем токен бота из переменной окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Не удалось получить токен бота. Установите переменную окружения TELEGRAM_BOT_TOKEN.")
	}

	// Создаем нового бота с указанным токеном
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Устанавливаем отладочный режим (выводим все отправленные сообщения)
	bot.Debug = true

	log.Printf("Авторизован как %s", bot.Self.UserName)

	// Создаем канал для получения обновлений
	updates, err := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
	if err != nil {
		log.Panic(err)
	}

	// Обрабатываем полученные обновления
	for update := range updates {
		if update.Message == nil { // игнорируем обновления, не являющиеся сообщениями
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Отвечаем на полученное сообщение
		reply := "Получено сообщение: " + update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
