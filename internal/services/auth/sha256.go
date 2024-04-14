package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func encodeSHA256(password []byte) string {
	// Создаем новый объект хеша SHA-256
	sh := sha256.New()

	// Записываем данные в хеш
	sh.Write(password)

	// Вычисляем финальный хеш и преобразуем его в строку в шестнадцатеричном формате
	hashedPassword := sh.Sum(nil)
	hexHash := hex.EncodeToString(hashedPassword)

	return hexHash
}
