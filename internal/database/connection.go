package database

import (
	"database/sql"
	"fmt"
	"log"

	"auth/internal/config"

	_ "github.com/lib/pq"
)

// getConnStr возвращает строку подключения к базе данных PostgreSQL на основе конфигурации.
func getConnStr() string {
	// Получаем конфигурацию из файла.
	conf, _ := config.GetConfig()
	// Формируем строку подключения на основе параметров конфигурации.
	str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.DatabaseName, conf.Postgres.SSLMode, conf.Postgres.Host)
	return str
}

// connStr содержит строку подключения к базе данных PostgreSQL.
var connStr = getConnStr()

// DB открывает соединение с базой данных PostgreSQL и возвращает указатель на объект *sql.DB для выполнения запросов.
func DB() (*sql.DB, error) {
	// Открываем соединение с базой данных PostgreSQL, используя строку подключения connStr.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// Если произошла ошибка при открытии соединения, записываем ошибку в журнал и возвращаем nil и ошибку.
		log.Fatal(err)
		return nil, err
	}

	// Возвращаем указатель на объект *sql.DB и nil ошибки, если соединение успешно открыто.
	return db, nil
}
