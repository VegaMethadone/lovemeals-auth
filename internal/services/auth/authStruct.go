package auth

// User представляет собой структуру данных для представления пользователей.
type User struct {
	Id    int    // Id - идентификатор пользователя.
	Phone string // Phone - телефонный номер пользователя.
}

// Executor представляет собой структуру данных для представления исполнителей.
type Executor struct {
	Id       int    // Id - идентификатор исполнителя.
	Login    string // Login - логин исполнителя.
	Password []byte // Password - пароль исполнителя (обычно в виде байтов для безопасности).
	Phone    string // Phone - телефонный номер исполнителя.
}
