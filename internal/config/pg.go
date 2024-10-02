package config

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

// Чтение кофига для postgresql
type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

// Инициализация конфига для бд
func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}