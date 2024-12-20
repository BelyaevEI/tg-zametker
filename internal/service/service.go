package service

import (
	"sync"

	"github.com/BelyaevEI/tg-zametker/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Имплементация сервисного слоя
type Servicer interface {
	Commands(update tgbotapi.Update) tgbotapi.MessageConfig
	NoteMenu(update tgbotapi.Update) tgbotapi.MessageConfig
	Start(update tgbotapi.Update) tgbotapi.MessageConfig
	NotFound(update tgbotapi.Update) tgbotapi.MessageConfig
	HandleText(update tgbotapi.Update) tgbotapi.MessageConfig
}

type serv struct {
	repository repository.Repositorer
	state      map[int64]string // Состояние для пользователей
	mu         sync.Mutex       // Мьютекс для безопасного доступа к состоянию
	numberNote int64            //Запрошенный номер при редактировании
}

func NewService(repository repository.Repositorer) Servicer {
	return &serv{
		repository: repository,
		state:      make(map[int64]string),
	}
}
