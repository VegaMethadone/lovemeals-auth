package database

import (
	"log"
)

// CreateNewUserDB создает новую запись пользователя в базе данных.
// Принимает телефонный номер нового пользователя.
// Возвращает ошибку, если произошла ошибка при добавлении записи в базу данных
// или при обновлении последнего идентификатора пользователей в метаданных.
func CreateNewUserDB(phone string) (int, error) {
	// Получаем соединение с базой данных
	db, err := DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return -1, err
	}

	// Получаем последний идентификатор пользователя
	newId, err := GetUsersOrExecutorsLastId(true)
	if err != nil {
		log.Println("Could not retrieve the last user ID")
		return -1, err
	}
	newId += 1 // Увеличиваем идентификатор для нового пользователя

	// Формируем SQL-запрос на добавление нового пользователя в базу данных
	query := "INSERT INTO users (user_id, phone) VALUES ($1, $2)"
	_, err = db.Exec(query, newId, phone)
	if err != nil {
		log.Println("An error occurred while adding a new user:", err)
		return -1, err
	}

	// Обновляем последний идентификатор пользователя в метаданных
	err = ChangeUsersOrExecutorsLastId(newId, true)
	if err != nil {
		log.Println("An error occurred while updating the last user ID:", err)
		return -1, err
	}

	// Возвращаем nil, если все операции выполнены успешно
	return newId, nil
}

// DeleteUser удаляет пользователя из базы данных по его идентификатору.
// Принимает идентификатор пользователя в качестве параметра.
// Возвращает ошибку, если произошла ошибка при удалении пользователя из базы данных.
func DeleteUserDB(id int) error {
	// Получаем соединение с базой данных
	db, err := DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return err
	}
	defer db.Close()

	// Формируем SQL-запрос на удаление пользователя по его идентификатору
	query := "DELETE FROM users WHERE user_id =$1"
	_, err = db.Exec(query, id)
	if err != nil {
		log.Println("An error occurred while deleting the user:", err)
		return err
	}

	// Возвращаем nil, если пользователь успешно удален
	return nil
}

// CheckUser проверяет наличие пользователя в базе данных по номеру телефона.
// Принимает номер телефона в качестве параметра.
// Возвращает идентификатор пользователя, если он найден, или -1 и ошибку, если пользователь не найден или произошла ошибка при выполнении запроса.
func CheckUserDB(phone string) (int, string, error) {
	// Получаем соединение с базой данных
	db, err := DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return -2, "", err
	}
	defer db.Close()

	// Формируем SQL-запрос на поиск пользователя по номеру телефона
	query := "SELECT user_id FROM users WHERE phone=$1"
	res, err := db.Query(query, phone)
	if err != nil {
		log.Println("Could not find users:", err)
		return -1, "", err
	}
	defer res.Close()

	var userID int
	// Проверяем, есть ли результаты запроса
	if res.Next() {
		// Считываем идентификатор пользователя из результата запроса
		err = res.Scan(&userID)
		if err != nil {
			log.Println("Failed to scan user id:", err)
			return -1, "", err
		}
	} else {
		// Если пользователя с указанным номером телефона не найдено
		log.Println("User not found with phone number:", phone)
		return -1, "", nil
	}
	// Возвращаем идентификатор пользователя и nil, если пользователь найден
	return userID, phone, nil
}

// UpdateUserData обновляет номер телефона пользователя в базе данных по указанному идентификатору.
// Функция принимает идентификатор пользователя и новый номер телефона в качестве аргументов.
// Если успешно устанавливается соединение с базой данных, функция выполняет запрос на обновление данных.
// После выполнения запроса соединение с базой данных закрывается.
// В случае ошибки при установлении соединения, выполнении запроса или закрытии соединения,
// функция возвращает ошибку. В противном случае возвращает nil.
func UpdateUserDataDB(id int, phone string) error {
	// Установка соединения с базой данных
	db, err := DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return err
	}

	defer db.Close() // Закрытие соединения с базой данных после завершения функции

	// SQL-запрос на обновление номера телефона пользователя по его идентификатору
	query := "UPDATE users SET phone=$1 WHERE id=$2"
	_, err = db.Exec(query, phone, id)
	if err != nil {
		log.Println("Failed to update number by ID:", phone, id)
		return err
	}
	return nil // Возврат nil в случае успешного обновления
}
