package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ExecutorJWT генерирует JWT (JSON Web Token) для исполнителя (executor) на основе предоставленных данных,
// таких как идентификатор, логин, номер телефона и секретный ключ.
// Возвращает строку с JWT в случае успеха или пустую строку в случае ошибки.
func ExecutorJWT(id int, login, phone, key string) string {
	// Создаем новый токен с использованием метода подписи HS256
	token := jwt.New(jwt.SigningMethodHS256)

	// Получаем доступ к утверждениям токена для добавления их значений
	claims := token.Claims.(jwt.MapClaims)

	// Устанавливаем утверждения токена
	claims["id"] = id                                     // Идентификатор исполнителя
	claims["login"] = login                               // Логин исполнителя
	claims["phone"] = phone                               // Номер телефона исполнителя
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Время истечения токена (24 часа)

	// Секретный ключ для подписи токена
	secretKey := key

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		// В случае ошибки подписи токена, выводим сообщение об ошибке в журнал и возвращаем пустую строку
		log.Println("Error signing executor token:", err)
		return ""
	}

	// Выводим в журнал сгенерированный JWT
	log.Println("Generated JWT:", tokenString)

	// Возвращаем сгенерированный JWT
	return tokenString
}

// ExecutorParseJWT анализирует переданный JWT (JSON Web Token) и извлекает данные исполнителя из него.
// Возвращает указатель на структуру Executor, содержащую извлеченные данные исполнителя, или nil в случае ошибки.
func ExecutorParseJWT(tokenString string) *Executor {
	// Парсим JWT, используя секретный ключ "ALAH"
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte("ALAH"), nil
	})

	// Обработка ошибок при парсинге JWT
	if err != nil {
		fmt.Println("Error parsing executor token:", err)
		return nil
	}

	// Создаем новый экземпляр структуры Executor для хранения извлеченных данных
	gotExecutor := &Executor{}

	// Проверяем, является ли токен действительным
	if token.Valid {
		// Извлекаем утверждения (claims) из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			// Преобразуем и извлекаем данные из утверждений
			gotExecutor.Id = int(claims["id"].(float64))   // Преобразуем float64 в int для идентификатора исполнителя
			gotExecutor.Phone = claims["phone"].(string)   // Получаем номер телефона исполнителя
			gotExecutor.Login = claims["login"].(string)   // Получаем логин исполнителя
			gotExecutor.ExpToken = claims["exp"].(float64) // Получаем время истечения токена

			// Выводим данные токена в журнал для отладки
			log.Println("TOKEN DATA", *gotExecutor)
		} else {
			// В случае ошибки получения утверждений выводим сообщение в журнал
			log.Println("Error getting token data")
		}
	} else {
		// В случае недействительного токена выводим сообщение в журнал
		log.Println("INVALID token", tokenString)
	}

	// Возвращаем указатель на структуру Executor с извлеченными данными или nil в случае ошибки
	return gotExecutor
}
