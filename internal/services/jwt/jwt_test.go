package jwt

import (
	"testing"
)

func TestUserJWTAndParse(t *testing.T) {
	// Подготовка данных для теста
	id := 1                // Идентификатор пользователя
	phone := "79657775481" // Номер телефона пользователя
	key := "ALAH"          // Ключ для создания JWT

	// Генерация JWT
	tokenString, err := UserJWT(id, phone, key)
	if err != nil {
		t.Errorf("%v", err)
	}

	// Проверка корректности генерации JWT
	if tokenString == "" {
		t.Error("Expected non-empty token string, got empty")
	}

	// Парсинг JWT
	parsedToken := UserParseJWT(tokenString)

	// Проверка данных, извлеченных из JWT
	if parsedToken.Id != id {
		t.Errorf("Expected id %d, got %d", id, parsedToken.Id)
	}

	if parsedToken.Phone != phone {
		t.Errorf("Expected phone %s, got %s", phone, parsedToken.Phone)
	}

}

// TestExecutorJWTandParse тестирует функции ExecutorJWT и ExecutorParseJWT.
func TestExecutorJWTandParse(t *testing.T) {
	// Подготовка данных для теста
	id := 2                    // Идентификатор пользователя
	login := "swine@gmail.com" // Логин пользователя
	phone := "79657775481"     // Номер телефона пользователя
	key := "ALAH"              // Ключ для создания JWT

	// Генерация JWT
	tokenString := ExecutorJWT(id, login, phone, key)

	// Проверка корректности генерации JWT
	if tokenString == "" {
		t.Error("Ожидалась непустая строка токена, получена пустая")
	}

	// Парсинг JWT
	parsedToken := ExecutorParseJWT(tokenString)

	// Проверка данных, извлеченных из JWT
	if parsedToken.Id != id {
		t.Errorf("Ожидался id %d, получен %d", id, parsedToken.Id)
	}

	if parsedToken.Login != login {
		t.Errorf("Ожидался логин %s, получен %s", login, parsedToken.Login)
	}

	if parsedToken.Phone != phone {
		t.Errorf("Ожидался номер телефона %s, получен %s", phone, parsedToken.Phone)
	}
}
