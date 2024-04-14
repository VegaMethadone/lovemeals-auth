package auth

import (
	"errors"
	"regexp"
	"unicode"
)

func validateNumber(phone string) error {
	// Проверяем длину телефонного номера
	if len(phone) != 11 {
		return errors.New("invalid phone number: len(phone) != 11")
	}

	// Проверяем, начинается ли телефонный номер с '7' или '8'
	if phone[0] != '7' && phone[0] != '8' {
		return errors.New("invalid phone number: begins with not 7 or 8")
	}

	// Проверяем, содержит ли телефонный номер только цифры
	for _, char := range phone {
		if !unicode.IsDigit(char) {
			return errors.New("invalid phone number: contains letters or symbols")
		}
	}
	return nil
}

func validateEmail(email string) bool {
	// Регулярное выражение для проверки корректности email
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Проверка email с использованием регулярного выражения
	match, _ := regexp.MatchString(regex, email)

	return match
}
