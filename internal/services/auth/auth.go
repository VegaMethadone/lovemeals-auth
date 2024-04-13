package auth

import (
	"errors"
	"unicode"
)

type authAPI struct {
	User     UserStorage
	Executor ExecutorStorage
}

// UserStorage представляет собой интерфейс для хранения операций с пользователями.
type UserStorage interface {
	// CreateUser создает нового пользователя в хранилище и возвращает его идентификатор или ошибку, если что-то пошло не так.
	CreateUser(phone string) (int, error)

	// GetUser возвращает идентификатор пользователя из хранилища или ошибку, если что-то пошло не так.
	GetUser() (int, error)

	// Logout осуществляет выход пользователя из системы по токену и возвращает true, если процесс завершился успешно, или ошибку, если что-то пошло не так.
	Logout(token string) (bool, error)
}

// ExecutorStorage представляет собой интерфейс для хранения операций с исполнителями.
type ExecutorStorage interface {
	// CreateExecutor создает нового исполнителя в хранилище с указанным логином, паролем и телефонным номером, возвращая true в случае успеха или ошибку, если что-то пошло не так.
	CreateExecutor(login string, password []byte, phone string) (int, error)

	// GetExecutor возвращает идентификатор, логин и телефонный номер исполнителя из хранилища, используя указанный логин и пароль, или ошибку, если что-то пошло не так.
	GetExecutor(login string, password []byte) (int, string, string)

	// Logout осуществляет выход исполнителя из системы по токену и возвращает true, если процесс завершился успешно, или ошибку, если что-то пошло не так.
	Logout(token string) (bool, error)
}

func (u *User) CreateUser(phone string) (int, error) {
	if len(phone) != 11 {
		return -1, errors.New("Invalid phone number: len(phone) != 11")
	}
	if phone[0] != '7' && phone[0] != '8' {
		return -1, errors.New("Invalid phone number: begins with not 7 or 8")
	}
	for _, char := range phone {
		if !unicode.IsDigit(char) {
			return -1, errors.New("Invalid phone number: contains letters or symbols")
		}
	}

}
