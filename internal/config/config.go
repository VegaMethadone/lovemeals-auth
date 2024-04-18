package config

import (
	"encoding/json"
	"log"
	"os"
)

func GetConfig() (*Config, error) {
	// Чтение содержимого файла конфигурации.
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	// Создание экземпляра конфигурации для заполнения данными из файла.
	conf := &Config{}

	// Разбор содержимого файла JSON и заполнение структуры конфигурации.
	err = json.Unmarshal(file, conf)
	if err != nil {
		// Если произошла ошибка при разборе JSON, выводим сообщение и завершаем программу.
		log.Fatalf("Failed to parse config JSON: %v", err)
		return nil, err
	}

	// Возвращаем заполненную структуру конфигурации.
	return conf, nil
}
