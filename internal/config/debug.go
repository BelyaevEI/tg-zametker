package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	debugEnvName = "DEBUG"
)

// Чтение токена из переменной окружения
type DebugConfig interface {
	Mode() bool
}

type debugConfig struct {
	mode bool
}

// Инициализация конфига токена
func NewDebugModeConfig() (DebugConfig, error) {
	modeStr := os.Getenv(debugEnvName)
	if len(modeStr) == 0 {
		return nil, errors.New("debug mode is empty")
	}

	modeBool, err := strconv.ParseBool(modeStr)
	if err != nil {
		return nil, errors.New("convert debug mode is failed")
	}

	return &debugConfig{
		mode: modeBool,
	}, nil
}

func (cfg *debugConfig) Mode() bool {
	return cfg.mode
}
