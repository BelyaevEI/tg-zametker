package service

import (
	"github.com/BelyaevEI/tg-zametker/internal/repository"
)

// Имплементация сервисного слоя
type Servicer interface{}

type serv struct {
	repository repository.Repositorer
}

func NewService(repository repository.Repositorer) Servicer {
	return &serv{
		repository: repository,
	}
}
