package config

import (
	"errors"
	"os"
)

const (
	tokenEnvName = "TOKEN"
)

// Чтение токена из переменной окружения
type TokenConfig interface {
	Token() string
}

type tokenConfig struct {
	token string
}

// Инициализация конфига токена
func NewTokenConfig() (TokenConfig, error) {
	token := os.Getenv(tokenEnvName)
	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}

	return &tokenConfig{
		token: token,
	}, nil
}

func (cfg *tokenConfig) Token() string {
	return cfg.token
}
