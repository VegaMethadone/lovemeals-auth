package auth

import (
	"auth/internal/database"
	"auth/internal/services/jwt"
	"errors"
	"log"
)

// CreateUser создает нового пользователя с заданным телефонным номером
func (u *user) NewUser(phone string) (int, error) {
	// Проверяем валидность номера телефона
	if err := validateNumber(phone); err != nil {
		return -1, err
	}

	// Создаем нового пользователя в базе данных, используя заданный телефонный номер
	id, err := database.CreateNewUserDB(phone)
	if err != nil {
		// Если возникла ошибка при создании пользователя в базе данных, возвращаем ошибку
		log.Println("Fail to create new user:", err)
		return id, err
	}

	// Возвращаем идентификатор нового пользователя
	return id, nil
}

// NewExecutor создает нового исполнителя в системе.
// Он принимает адрес электронной почты, пароль в виде среза байтов и номер телефона в качестве параметров.
// Функция проверяет корректность формата номера телефона и адреса электронной почты.
// Пароль хешируется с использованием алгоритма SHA256.
// Затем создается новый исполнитель в базе данных с переданными данными.
// Если все операции завершаются успешно, функция возвращает идентификатор нового исполнителя и nil в качестве ошибки.
// Если возникает ошибка при выполнении какого-либо шага, функция возвращает -1 в качестве идентификатора исполнителя
// и ошибку, описывающую причину сбоя операции.
func (e *executor) NewExecutor(login string, password []byte, phone string) (int, error) {
	// Проверяем корректность формата номера телефона.
	if err := validateNumber(phone); err != nil {
		return -1, err // Возвращаем -1 и ошибку, если номер телефона некорректен
	}

	// Проверяем корректность формата электронной почты.
	if valid := validateEmail(login); !valid {
		return -1, errors.New("invalid mail format") // Возвращаем -1 и ошибку, если email некорректен
	}

	// Хешируем пароль с использованием SHA256.
	encodedPassword := encodeSHA256(password)
	if len(encodedPassword) != 64 {
		return -1, errors.New("failed to encode password") // Возвращаем -1 и ошибку, если хеширование пароля не удалось
	}

	// Создаем нового исполнителя в базе данных.
	gotID, err := database.CreateNewExecutorDB(login, encodedPassword, phone)
	if err != nil {
		return gotID, err // Возвращаем идентификатор и ошибку, если возникла ошибка при создании исполнителя в базе данных
	}

	// Все операции завершены успешно. Возвращаем идентификатор нового исполнителя.
	return gotID, nil
}

// GetUser получает телефонный номер пользователя и возвращает JWT токен,
// а также проверяет наличие пользователя в базе данных и валидность номера телефона.
func (u *user) GetUser(phone string) (string, error) {
	// Проверяем валидность номера телефона
	if err := validateNumber(phone); err != nil {
		return "", err
	}

	// Проверяем наличие пользователя в базе данных
	userId, userPhone, err := database.CheckUserDB(phone)
	if err != nil {
		return "", err
	}

	// Создаем JWT токен для пользователя
	token, err := jwt.UserJWT(userId, userPhone, keyWord)
	if err != nil {
		log.Println("Failed to create JWT token:", err)
		return "", err
	}

	return token, nil
}

// GetExecutor осуществляет аутентификацию исполнителя по его адресу электронной почты и паролю.
// Сначала функция проверяет корректность формата адреса электронной почты.
// Затем пароль хешируется с использованием алгоритма SHA256.
// После этого происходит проверка наличия исполнителя в базе данных с использованием введенных учетных данных.
// Если исполнитель найден и учетные данные верны, функция генерирует JWT токен для аутентификации исполнителя.
// Токен возвращается в случае успешной аутентификации, а также nil в качестве ошибки.
// Если возникает ошибка при любом из шагов, функция возвращает пустую строку и описание ошибки.
func (e *executor) GetExecutor(login string, password []byte) (string, error) {
	// Проверяем корректность формата электронной почты.
	if valid := validateEmail(login); !valid {
		return "", errors.New("invalid mail format")
	}

	// Хешируем пароль с использованием SHA256.
	encodedPassword := encodeSHA256(password)
	if len(encodedPassword) != 64 {
		return "", errors.New("failed to encode password")
	}

	// Проверяем наличие исполнителя в базе данных с использованием введенных учетных данных.
	gotID, gotLogin, gotPhone, err := database.CheckExecutorDB(login, encodedPassword)
	if err != nil {
		return "", err // Возвращаем ошибку, если произошла ошибка при проверке учетных данных.
	}

	// Генерируем JWT токен для аутентификации исполнителя.
	token := jwt.ExecutorJWT(gotID, gotLogin, gotPhone, keyWord)
	if token == "" {
		return "", errors.New("failed to generate token") // Возвращаем ошибку, если не удалось сгенерировать токен.
	}

	// Возвращаем токен и nil в случае успеха.
	return token, nil
}

// Logout завершает сеанс пользователя, проверяя валидность переданного токена.
func (u *user) Logout(token string) (string, error) {
	// Проверяем валидность токена, вызывая функцию UserParseJWT из пакета jwt
	if gotUser := jwt.UserParseJWT(token); gotUser != nil {
		// Если токен недействителен, возвращаем ошибку
		return "", errors.New("invalid token")
	}
	// Если токен действителен, возвращаем пустую строку в качестве токена и nil в качестве ошибки
	return "", nil
}

// Logout завершает сеанс пользователя, проверяя валидность переданного токена.
func (e *executor) Logout(token string) (string, error) {
	// Проверяем валидность токена, вызывая функцию UserParseJWT из пакета jwt
	if gotUser := jwt.UserParseJWT(token); gotUser != nil {
		// Если токен недействителен, возвращаем ошибку
		return "", errors.New("invalid token")
	}
	// Если токен действителен, возвращаем пустую строку в качестве токена и nil в качестве ошибки
	return "", nil
}
