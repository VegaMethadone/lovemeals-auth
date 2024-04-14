package database

import (
	"fmt"
	"testing"
)

// TestCreateNewExecutorDB тестирует функцию CreateNewExecutorDB, которая добавляет нового исполнителя (executor) в базу данных.
// Тест создает нового исполнителя с указанным логином, паролем и номером телефона.
// После добавления исполнителя в базу данных тест проверяет, успешно ли он был добавлен.
// Если происходит ошибка при добавлении исполнителя, тест завершается с фатальной ошибкой.
func TestCreateNewExecutorDB(t *testing.T) {
	// Данные нового исполнителя для тестирования
	login := "shpack@gmail.com"
	password := "3fafadb34c8637d470635816200651248504576fe388702ed605739cdd713c56" // Примерный хэш пароля
	phone := "89302214018"

	// Добавление нового исполнителя в базу данных
	id, err := CreateNewExecutorDB(login, password, phone)
	if err != nil {
		t.Fatal("Can not to add new executor to DataBase:", err)
	} else {
		fmt.Println("Executor created with ID", id)
	}
}

// TestCheckExecutor тестирует функцию CheckExecutorDB, которая проверяет наличие исполнителя в базе данных по логину и паролю.
// Функция создает нового исполнителя в базе данных с заданным логином, паролем и номером телефона.
// Затем она вызывает функцию CheckExecutorDB для проверки наличия созданного исполнителя в базе данных.
// Если исполнитель не найден, тест завершается с ошибкой.
// Если исполнитель найден, тест проверяет, что полученные данные соответствуют ожидаемым.
// Если данные не соответствуют ожидаемым, тест также завершается с ошибкой.
// В конце тест удаляет созданного исполнителя из базы данных.
func TestCheckExecutor(t *testing.T) {
	// Задаем данные исполнителя для тестирования
	login := "Jop@gmail.com"                   // Логин исполнителя
	password := "c29rc20ru30vmu20v230mvr20r2e" // Пароль исполнителя
	phone := "89217361731"                     // Номер телефона исполнителя

	// Создаем нового исполнителя в базе данных
	id, err := CreateNewExecutorDB(login, password, phone) // Вызываем функцию для создания нового исполнителя
	if err != nil {
		t.Fatal("Can not to add new executor to DataBase:", err) // Если произошла ошибка при создании исполнителя, завершаем тест
	}
	fmt.Println("EXECUTOR CREATED WITH ID & LOGIN:", id, login) // Выводим сообщение о создании исполнителя

	// Проверяем наличие исполнителя в базе данных
	gotId, gotLogin, gotPhone, err := CheckExecutorDB(login, password) // Вызываем функцию для проверки наличия исполнителя
	if err != nil {
		t.Fatalf("Can not to find executor with login: %s -->  %v", login, err) // Если исполнитель не найден, завершаем тест с ошибкой
	}
	fmt.Println("FOUND EXECUTOR WITH ID BY LOGIN:", gotId, gotLogin) // Выводим сообщение о нахождении исполнителя в базе данных

	// Проверяем, что полученные данные соответствуют ожидаемым
	if login != gotLogin || phone != gotPhone {
		t.Fatalf("Expected: %s %s \n Got: %s %s", login, password, phone, gotPhone) // Если данные не соответствуют ожидаемым, завершаем тест с ошибкой
	}

	// Удаляем исполнителя из базы данных
	err = DeleteExecutorDB(gotId) // Вызываем функцию для удаления исполнителя
	if err != nil {
		t.Fatal("An error occurred while deleting the executor:", err) // Если произошла ошибка при удалении исполнителя, завершаем тест
	}
	fmt.Println("EXECUTOR IS DELETED WITH ID:", gotId) // Выводим сообщение об успешном удалении исполнителя
}
