package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// UserJWT генерирует JWT (JSON Web Token) на основе предоставленного идентификатора пользователя,
// номера телефона и секретного ключа. Возвращает строку с JWT в случае успеха или пустую строку в случае ошибки.
func UserJWT(id int, phone, key string) (string, error) {
	// Создаем новый токен с использованием метода подписи HS256
	token := jwt.New(jwt.SigningMethodHS256)

	// Получаем доступ к утверждениям токена для добавления их значений
	claims := token.Claims.(jwt.MapClaims)

	// Устанавливаем утверждения токена
	claims["id"] = id                                     // Идентификатор пользователя
	claims["phone"] = phone                               // Номер телефона пользователя
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Время истечения токена (24 часа)

	// Секретный ключ для подписи токена
	secretKey := key

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		// В случае ошибки подписи токена, выводим сообщение об ошибке в журнал и возвращаем пустую строку
		log.Println("Error signing user token:", err)
		return "", err
	}

	// Выводим в консоль сгенерированный JWT
	log.Println("Generated JWT:", tokenString)

	// Возвращаем сгенерированный JWT
	return tokenString, nil
}

// UserParseJWT анализирует переданный JWT (JSON Web Token) и извлекает данные пользователя из него.
// Возвращает указатель на структуру User, содержащую извлеченные данные пользователя, или nil в случае ошибки.
func UserParseJWT(tokenString string) *User {
	// Парсим JWT, используя секретный ключ "ALAH"
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("ALAH"), nil
	})

	// Обработка ошибок при парсинге JWT
	if err != nil {
		fmt.Println("Error parsing user token:", err)
		return nil
	}

	// Создаем новый экземпляр структуры User для хранения извлеченных данных
	gotUser := &User{}

	// Проверяем, является ли токен действительным
	if token.Valid {
		// Извлекаем утверждения (claims) из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			// Преобразуем и извлекаем данные из утверждений
			gotUser.Id = int(claims["id"].(float64))   // Преобразуем float64 в int для идентификатора пользователя
			gotUser.Phone = claims["phone"].(string)   // Получаем номер телефона пользователя
			gotUser.ExpToken = claims["exp"].(float64) // Получаем время истечения токена

			// Выводим данные токена в журнал для отладки
			log.Println("TOKEN DATA", *gotUser)
		} else {
			// В случае ошибки получения утверждений выводим сообщение в журнал
			log.Println("Error getting token data")
		}
	} else {
		// В случае недействительного токена выводим сообщение в журнал
		log.Println("INVALID token", tokenString)
	}

	// Возвращаем указатель на структуру User с извлеченными данными или nil в случае ошибки
	return gotUser
}
