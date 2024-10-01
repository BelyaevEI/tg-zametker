package config

import "github.com/joho/godotenv"

// Загружаем файл в переменные окружения
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
