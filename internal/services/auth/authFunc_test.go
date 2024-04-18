package auth

import (
	"auth/internal/database"
	"log"
	"testing"
)

// Интеграционный тест для набора функций, относящихся к пользователю.
func TestUser(t *testing.T) {
	// Устанавливаем номер телефона для нового пользователя
	phone := "89007775540"
	// Создаем новый экземпляр API
	api := NewAPI()

	// Шаг 1: Регистрация нового пользователя
	id, err := api.User.NewUser(phone)
	if err != nil {
		t.Errorf("Failed to add new user with phone: %s", phone)
	}

	// Выводим подтверждение успешной регистрации пользователя
	log.Println("Registered new user with ID:", id)

	// Получение токена пользователя
	token, err := api.User.GetUser(phone)
	if err != nil {
		t.Errorf("Failed to get user: %v", err)
	}

	// Удаление токена пользователя (выход из системы)
	token, err = api.User.Logout(token)
	if err != nil {
		t.Errorf("Failed to delete token: %v", err)
	} else {
		// Выводим подтверждение успешного удаления токена
		log.Println("Token deleted", token)
	}

	// Удаление пользователя из базы данных
	if err := database.DeleteUserDB(id); err != nil {
		t.Errorf("Failed to delete user from db: %v", err)
	} else {
		// Выводим подтверждение успешного удаления пользователя из базы данных
		log.Println("User deleted with ID:", id)
	}
}

// Интеграционный тест для набора функций, относящихся к исполнителю.
func TestExecutor(t *testing.T) {
	// Задаем данные для нового исполнителя
	login := "swine@gmail.com"
	password := "qweqwe2077"
	phone := "89007775546"

	// Создаем новый экземпляр API
	api := NewAPI()

	// Создание нового исполнителя
	id, err := api.Executor.NewExecutor(login, []byte(password), phone)
	if err != nil {
		t.Errorf("Failed to create new executor with login: %s", login)
	}

	// Выводим подтверждение успешной регистрации исполнителя
	log.Println("Registered new executor with ID:", id)

	// Получение токена исполнителя
	token, err := api.Executor.GetExecutor(login, []byte(password))
	if err != nil {
		t.Errorf("Failed to get executor: %v", err)
	}

	// Удаление токена исполнителя (выход из системы)
	token, err = api.Executor.Logout(token)
	if err != nil {
		t.Errorf("Failed to delete token: %v", err)
	} else {
		// Выводим подтверждение успешного удаления токена
		log.Println("Token deleted", token)
	}

	// Удаление исполнителя из базы данных
	if err := database.DeleteExecutorDB(id); err != nil {
		t.Errorf("Failed to delete executor from db: %v", err)
	} else {
		// Выводим подтверждение успешного удаления исполнителя из базы данных
		log.Println("Executor deleted with ID:", id)
	}
}
