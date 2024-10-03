package app

import (
	"context"
	"log"

	"github.com/BelyaevEI/tg-zametker/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type App struct {
	bot             *tgbotapi.BotAPI
	serviceProvider *serviceProvider
}

// Создание структуры приложения
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Инициализация всех зависимостей приложения
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initBot,
		a.initService,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load("../config.env")
	if err != nil {
		return err
	}

	return nil
}

// Инициализация сервис провайдера
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

// Создаем нового бота с указанным токеном
func (a *App) initBot(ctx context.Context) error {

	bot, err := tgbotapi.NewBotAPI(a.serviceProvider.TokenConfig().Token())
	if err != nil {
		log.Fatalf("create bot is failed: %s", err.Error())
		return err
	}

	// Устанавливаем отладочный режим (выводим все отправленные сообщения)
	bot.Debug = a.serviceProvider.DebugConfig().Mode()

	a.bot = bot

	return nil
}

func (a *App) initService(ctx context.Context) error {
	a.serviceProvider.ImplementationApp(ctx)

	return nil
}

// Запуск бота
func (a *App) Run(ctx context.Context) error {

	log.Println("Bot is running...")

	// Получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := a.bot.GetUpdatesChan(u)
	if err != nil {
		return nil
	}

	for update := range updates {
		if update.Message == nil { // проверяем, что сообщение не пустое
			continue
		}

		if update.Message.IsCommand() { // если пришла команда
			switch update.Message.Command() {
			case "start":
				// Создаем reply-кнопки
				replyKeyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Кнопка 1"),
						tgbotapi.NewKeyboardButton("Кнопка 2"),
					),
				)

				// Отправляем сообщение с reply-кнопками
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите опцию:")
				msg.ReplyMarkup = replyKeyboard

				a.bot.Send(msg)
			}
		}

		switch update.Message.Text {
		case "Кнопка 1":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали Кнопка 1")
			a.bot.Send(msg)
		case "Кнопка 2":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали Кнопка 2")
			a.bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
			a.bot.Send(msg)
		}
	}

	return nil
}
