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
	bot.Debug = true

	a.bot = bot

	return nil
}

func (a *App) Run(ctx context.Context) error {

	log.Println("Bot is running...")

	// Получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5

	updates, err := a.bot.GetUpdatesChan(u)
	if err != nil {
		return nil
	}

	for update := range updates {
		if update.Message != nil { // Проверка, если пришло сообщение
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Ответ на сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName+"!")
			a.bot.Send(msg)
		}
	}

	return nil
}
