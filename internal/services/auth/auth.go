package auth

const keyWord string = "ALAH"

type user struct{}
type executor struct{}

// authAPI представляет API для аутентификации и управления пользователями и исполнителями.
type authAPI struct {
	User     UserStorage     // Хранилище операций с пользователями.
	Executor ExecutorStorage // Хранилище операций с исполнителями.
}

// UserStorage представляет собой интерфейс для хранения операций с пользователями.
type UserStorage interface {
	// NewUser создает нового пользователя в хранилище и возвращает его идентификатор или ошибку, если что-то пошло не так.
	NewUser(phone string) (int, error)

	// GetUser возвращает токен пользователя или ошибку, если что-то пошло не так.
	GetUser(phone string) (string, error)

	// Logout осуществляет выход пользователя из системы по токену и возвращает true, если процесс завершился успешно, или ошибку, если что-то пошло не так.
	Logout(token string) (string, error)
}

// ExecutorStorage представляет собой интерфейс для хранения операций с исполнителями.
type ExecutorStorage interface {
	// NewExecutor создает нового исполнителя в хранилище с указанным логином, паролем и телефонным номером, возвращая идентификатор исполнителя или ошибку, если что-то пошло не так.
	NewExecutor(login string, password []byte, phone string) (int, error)

	// GetExecutor возвращает идентификатор, логин и телефонный номер исполнителя из хранилища, используя указанный логин и пароль, или ошибку, если что-то пошло не так.
	GetExecutor(login string, password []byte) (string, error)

	// Logout осуществляет выход исполнителя из системы по токену и возвращает true, если процесс завершился успешно, или ошибку, если что-то пошло не так.
	Logout(token string) (string, error)
}

// NewAPI создает новый экземпляр authAPI с инициализированными хранилищами пользователей и исполнителей.
func NewAPI() *authAPI {
	return &authAPI{
		User:     &user{},     // Инициализируем хранилище пользователей.
		Executor: &executor{}, // Инициализируем хранилище исполнителей.
	}
}
