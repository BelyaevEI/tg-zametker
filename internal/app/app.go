package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/BelyaevEI/tg-zametker/internal/config"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	// Gracefull shutdown
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	ctx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Запускаем сам бот
	go func() {
		defer wg.Done()

		err := a.UpdatesFromTGServer(ctx)
		if err != nil {
			log.Fatalf("get updates is failed: %s", err.Error())
		}

	}()

	gracefulShutdown(ctx, cancel, wg)
	return nil
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-waitSignal():
		log.Println("terminating: via signal")
	}

	cancel()
	if wg != nil {
		wg.Wait()
	}
}

func waitSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}
