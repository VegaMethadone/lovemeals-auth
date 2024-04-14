package database

import (
	"fmt"
	"testing"
)

// TestCreateNewUserDB тестирует функцию CreateNewUserDB, которая добавляет нового пользователя в базу данных.
// Тест создает нового пользователя с указанным номером телефона и проверяет, был ли пользователь успешно добавлен в базу данных.
// Если пользователь успешно добавлен, тест завершается без ошибок, иначе завершается с фатальной ошибкой.
func TestCreateNewUserDB(t *testing.T) {
	// Номер телефона нового пользователя для тестирования
	phone := "77775553535"

	// Добавление нового пользователя в базу данных
	id, err := CreateNewUserDB(phone)
	if err != nil {
		t.Fatal("Can not to add new user to DataBase:", err)
	} else {
		fmt.Println("User created with ID:", id)
	}
}

// TestCheckUser тестирует функцию CheckUser, которая проверяет наличие пользователя в базе данных по номеру телефона.
// Тест создает нового пользователя с указанным номером телефона, затем выполняет запрос на проверку его наличия в базе данных.
// Если пользователь найден, тест удаляет его из базы данных.
// Если происходит ошибка при добавлении, поиске или удалении пользователя, тест завершается с фатальной ошибкой.
func TestCheckUser(t *testing.T) {
	// Номер телефона пользователя для тестирования
	phone := "89998887766"

	// Создание нового пользователя с указанным номером телефона
	id, err := CreateNewUserDB(phone)
	if err != nil {
		t.Fatal("Can not to add new user to DataBase:", err)
	}
	fmt.Println("USER CREATED WITH ID & PHONE:", id, phone)

	// Проверка наличия пользователя в базе данных по номеру телефона
	id, gotPhone, err := CheckUserDB(phone)
	if err != nil {
		t.Fatalf("Can not to find user with number: %s -->  %v", phone, err)
	}
	fmt.Println("FOUND USER WITH ID  BY PHONE:", id, gotPhone)

	// Удаление пользователя из базы данных
	err = DeleteUserDB(id)
	if err != nil {
		t.Fatal("An error occurred while deleting the user:", err)
	}
	fmt.Println("USER IS DELETED WITH ID:", id)
}

// TestUpdateUserData тестирует функцию UpdateUserData, которая обновляет номер телефона пользователя в базе данных.
// Тест создает нового пользователя с определенным номером телефона, затем находит его в базе данных и обновляет номер телефона.
// После успешного обновления телефона тест удаляет пользователя из базы данных.
// Если происходит ошибка при добавлении, поиске, обновлении или удалении пользователя, тест завершается с фатальной ошибкой.
func TestUpdateUserData(t *testing.T) {
	// Создание нового пользователя в базе данных с указанным номером телефона
	phone := "89998887766"
	id, err := CreateNewUserDB(phone)
	if err != nil {
		t.Fatal("Can not to add new user to DataBase:", err)
	}
	fmt.Println("USER CREATED WITH ID & PHONE:", id, phone)

	// Поиск пользователя в базе данных по номеру телефона
	id, gotPhone, err := CheckUserDB(phone)
	if err != nil {
		t.Fatalf("Can not to find user with number: %s -->  %v", phone, err)
	}
	fmt.Println("FOUND USER WITH ID  BY PHONE:", id, gotPhone)

	// Обновление номера телефона пользователя
	newPhone := "79653216699"
	err = UpdateUserDataDB(id, newPhone)
	if err != nil {
		t.Fatalf("Failed to update user phone %s  by ID: %d", newPhone, id)
	}
	fmt.Printf("USER WITH ID: %d IS UPDATED: old phone - %s, new phone - %s\n", id, gotPhone, newPhone)

	// Удаление пользователя из базы данных
	err = DeleteUserDB(id)
	if err != nil {
		t.Fatal("An error occurred while deleting the user:", err)
	}
	fmt.Println("USER IS DELETED WITH ID:", id)
}
