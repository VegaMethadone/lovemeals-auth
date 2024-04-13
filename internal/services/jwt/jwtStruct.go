package jwt

// User представляет собой структуру данных для представления пользователей.
type User struct {
	Id       int     // Id - идентификатор пользователя.
	Phone    string  // Phone - телефонный номер пользователя.
	ExpToken float64 // ExpToken - индификатор жизни токен
}

// Executor представляет собой структуру данных для представления исполнителей.
type Executor struct {
	Id       int     // Id - идентификатор исполнителя.
	Login    string  // Login - логин исполнителя.
	Phone    string  // Phone - телефонный номер исполнителя.
	ExpToken float64 // ExpToken - индификатор жизни токен
}
