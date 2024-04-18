package config

import (
	"log"
	"testing"
)

// TestGetConfig проверяет функцию GetConfig, которая загружает конфигурацию из файла.
func TestGetConfig(t *testing.T) {
	// Вызов функции GetConfig для загрузки конфигурации.
	conf, err := GetConfig()

	// Проверка наличия ошибки при загрузке конфигурации.
	if err != nil {
		t.Errorf("Failed to load config file: %v", err)
	} else {
		// Если загрузка прошла успешно, выводим значения конфигурации в консоль для отладки.
		log.Println("Config Version:", conf.Version)
		log.Println("Config Network:", conf.Network)
		log.Println("Config Postgres:", conf.Postgres)
	}
}
