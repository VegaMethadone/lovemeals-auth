package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// const connStr string = "user=postgres password=0000 dbname=testDB sslmode=disable host=database"
const connStr string = "user=postgres password=0000 dbname=testDB sslmode=disable"

// DB открывает соединение с базой данных PostgreSQL и возвращает указатель на объект *sql.DB для выполнения запросов.
func DB() (*sql.DB, error) {
	// Открываем соединение с базой данных PostgreSQL, используя строку подключения connStr
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// Если произошла ошибка при открытии соединения, записываем ошибку в журнал и возвращаем nil и ошибку
		log.Fatal(err)
		return nil, err
	}

	// Возвращаем указатель на объект *sql.DB и nil ошибки, если соединение успешно открыто
	return db, nil
}
