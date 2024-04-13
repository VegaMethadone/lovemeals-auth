package database

import (
	"auth/internal/database"
	"log"
)

// CreateNewExecutorDB добавляет нового исполнителя (executor) в базу данных.
// Функция принимает логин, пароль и номер телефона нового исполнителя.
// Сначала она получает последний идентификатор исполнителя из базы данных.
// Затем устанавливается соединение с базой данных и выполняется запрос на вставку данных в таблицу executors.
// Если вставка данных проходит успешно, идентификатор исполнителя обновляется, и изменения сохраняются в базе данных.
// В случае ошибки при получении последнего идентификатора, установке соединения, выполнении запроса, обновлении идентификатора
// или закрытии соединения с базой данных, функция возвращает ошибку. В противном случае возвращает nil.
func CreateNewExecutorDB(login, password, phone string) (int, error) {
	// Получение последнего идентификатора исполнителя из базы данных
	newId, err := database.GetUsersOrExecutorsLastId(false)
	if err != nil {
		log.Println("Failed to get last executor ID:", err)
		return -1, err
	}

	// Установка соединения с базой данных
	db, err := database.DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return -1, err
	}
	newId += 1

	// SQL-запрос на вставку данных нового исполнителя в таблицу executors
	query := "INSERT INTO executors (executor_id, login, password, phone) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, newId, login, password, phone)
	if err != nil {
		log.Println("Failed to add new executor to DataBase with data:", newId, login, password, phone)
		return -1, err
	}

	// Обновление последнего идентификатора исполнителя в базе данных
	err = database.ChangeUsersOrExecutorsLastId(newId, false)
	if err != nil {
		log.Println("An error occurred while updating the last executor ID:", err)
		return -1, err
	}

	// Закрытие соединения с базой данных
	defer db.Close()

	return newId, nil
}

// DeleteExecutor удаляет исполнителя из базы данных по указанному идентификатору.
// Функция принимает идентификатор исполнителя в качестве аргумента.
// Сначала она устанавливает соединение с базой данных.
// Затем выполняется SQL-запрос на удаление исполнителя из таблицы executors.
// После успешного выполнения запроса соединение с базой данных закрывается.
// В случае ошибки при установке соединения, выполнении запроса или закрытии соединения,
// функция возвращает ошибку. В противном случае возвращает nil.
func DeleteExecutorDB(id int) error {
	// Установка соединения с базой данных
	db, err := database.DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return err
	}
	defer db.Close() // Закрытие соединения с базой данных после завершения функции

	// SQL-запрос на удаление исполнителя из таблицы executors по его идентификатору
	query := "DELETE FROM executors WHERE executor_id=$1"
	_, err = db.Exec(query, id)
	if err != nil {
		log.Println("An error occurred while deleting the executor:", err)
		return err
	}

	return nil // Возврат nil в случае успешного удаления исполнителя
}

// CheckExecutor проверяет наличие исполнителя (executor) в базе данных по заданному логину и паролю.
// Функция принимает логин и пароль в качестве аргументов.
// Сначала она устанавливает соединение с базой данных.
// Затем выполняется SQL-запрос на выборку данных из таблицы executors по заданному логину.
// Если исполнитель найден, его данные сканируются и возвращаются в виде кортежа (id, login, phone).
// Если исполнитель не найден или указанный логин и пароль не совпадают с данными в базе данных, функция возвращает -1 и пустые строки.
// Если происходит ошибка при установке соединения, выполнении запроса или сканировании данных,
// функция возвращает соответствующую ошибку.
func CheckExecutorDB(login, password string) (int, string, string, error) {
	// Установка соединения с базой данных
	db, err := database.DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return -2, "", "", err
	}
	defer db.Close() // Закрытие соединения с базой данных после завершения функции

	// SQL-запрос на выборку данных исполнителя из таблицы executors по заданному логину
	query := "SELECT executor_id, login, password, phone FROM executors WHERE login=$1"
	res, err := db.Query(query, login)
	if err != nil {
		log.Println("Could not find executor", err)
		return -1, "", "", err
	}
	defer res.Close()

	var (
		gotId       int
		gotLogin    string
		gotPassword string
		gotPhone    string
	)
	if res.Next() {
		// Если исполнитель найден, его данные сканируются
		err = res.Scan(&gotId, &gotLogin, &gotPassword, &gotPhone)
		if err != nil {
			log.Println("Failed to scan executor id:", err)
			return -1, "", "", err
		}
	} else {
		// Если исполнитель не найден с заданным логином, возвращается ошибка
		log.Println("Executor not found with login:", login)
		return -1, "", "", nil
	}

	// Проверка соответствия логина и пароля
	if gotLogin != login || gotPassword != password {
		log.Println("Invalid login/password")
		return -1, "", "", nil
	}

	// Возвращение идентификатора, логина и номера телефона исполнителя
	return gotId, gotLogin, gotPhone, nil
}

// UpdateExecutorDB обновляет данные исполнителя в базе данных по указанному идентификатору.
// Функция принимает идентификатор исполнителя, а также новые значения логина, пароля и номера телефона.
// Сначала она устанавливает соединение с базой данных.
// Затем выполняется SQL-запрос на обновление данных исполнителя в таблице executors.
// После успешного выполнения запроса соединение с базой данных закрывается.
// В случае ошибки при установке соединения, выполнении запроса или закрытии соединения,
// функция возвращает ошибку. В противном случае возвращает nil.
func UpdateExecutorDB(id int, login, password, phone string) error {
	// Установка соединения с базой данных
	db, err := database.DB()
	if err != nil {
		log.Println("Failed to establish connection to the database")
		return err
	}
	defer db.Close() // Закрытие соединения с базой данных после завершения функции

	// SQL-запрос на обновление данных исполнителя в таблице executors по его идентификатору
	query := "UPDATE executors SET login=$1, password=$2, phone=$3 WHERE executor_id=$4"
	_, err = db.Exec(query, login, password, phone, id)
	if err != nil {
		log.Println("Failed to update executor by ID:", id)
		return err
	}

	return nil // Возврат nil в случае успешного обновления данных исполнителя
}
