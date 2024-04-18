package server

import (
	"auth/internal/config"
	"net/http"
	"time"
)

// NewServer создает новый HTTP-сервер на основе конфигурации и возвращает его.
func NewServer() *http.Server {
	// Получение конфигурации из файла.
	conf, err := config.GetConfig()
	if err != nil {
		// Если произошла ошибка при загрузке конфигурации, возвращаем nil.
		return nil
	}

	// Создание нового роутера для обработки запросов.
	router := NewRouter()

	// Формирование адреса сервера на основе конфигурации.
	adr := conf.Network.Address + conf.Network.Port

	// Создание HTTP-сервера с указанными параметрами.
	srv := &http.Server{
		// Установка созданного роутера как обработчика запросов.
		Handler: router,
		// Указание адреса и порта, на котором будет слушать сервер.
		Addr: adr,
		// Установка времени ожидания для операций записи.
		WriteTimeout: time.Duration(conf.Network.WriteTimeout) * time.Second,
		// Установка времени ожидания для операций чтения.
		ReadTimeout: time.Duration(conf.Network.ReadTimeout) * time.Second,
	}

	// Возвращаем созданный сервер.
	return srv
}
