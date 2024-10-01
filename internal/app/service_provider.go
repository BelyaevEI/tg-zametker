package app

import (
	"log"

	"github.com/BelyaevEI/tg-zametker/internal/config"
)

type serviceProvider struct {
	tokenConfig config.TokenConfig
}

// Создание сервис провайдера
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// Создание обьекта токена
func (s *serviceProvider) TokenConfig() config.TokenConfig {
	if s.tokenConfig == nil {
		cfg, err := config.NewTokenConfig()
		if err != nil {
			log.Fatalf("failed to get token config: %s", err.Error())
		}

		s.tokenConfig = cfg
	}

	return s.tokenConfig
}
