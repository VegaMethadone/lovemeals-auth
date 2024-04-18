package server

import (
	"auth/internal/services/auth"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var api = auth.NewAPI()

func CreateExecutorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func NewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что метод запроса POST.
		if r.Method != http.MethodPost {
			// Если метод не POST, отправляем клиенту статус "Метод не разрешен" и завершаем обработку запроса.
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Создаем новый экземпляр newUser для декодирования JSON-тела запроса.
		user := &newUser{}

		// Декодируем JSON-тело запроса в структуру newUser.
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			// Если произошла ошибка при декодировании, отправляем клиенту статус BadRequest и логируем ошибку.
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("Failed to decode request body: %v", err)
			return
		}

		// Создаем нового пользователя, используя метод NewUser из API.
		newID, err := api.User.NewUser(user.Phone)
		if err != nil {
			// Если произошла ошибка при создании пользователя, отправляем клиенту статус InternalServerError и логируем ошибку.
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Server error: %v", err)
			return
		}

		// Если создание пользователя прошло успешно, логируем ID нового пользователя и отправляем клиенту статус OK.
		log.Println("User created with ID:", newID)
		//  В дальнейшем придумать, как возвращать от сюда id и создавать профиль на другом сервере
		w.WriteHeader(http.StatusOK)
	}
}

func NewExecutor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что метод запроса POST.
		if r.Method != http.MethodPost {
			// Если метод не POST, отправляем исполнителю статус "Метод не разрешен" и завершаем обработку запроса.
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Создаем новый экземпляр newExecutor для декодирования JSON-тела запроса.
		newExecutor := &newExecutor{}

		// Декодируем JSON-тело запроса в структуру newExecutor.
		err := json.NewDecoder(r.Body).Decode(newExecutor)
		if err != nil {
			// Если произошла ошибка при декодировании, отправляем исполнителю статус BadRequest и логируем ошибку.
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("Failed to decode request body: %v", err)
			return
		}

		// Создаем нового исполнителя, используя метод NewExecutor из API.
		newID, err := api.Executor.NewExecutor(newExecutor.Login, []byte(newExecutor.Password), newExecutor.Phone)
		if err != nil {
			// Если произошла ошибка при создании исполнителя, отправляем исполнителю статус InternalServerError и логируем ошибку.
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Server error: %v", err)
			return
		}

		// Если создание исполнителя прошло успешно, логируем ID нового исполнителя и отправляем клиенту статус OK.
		log.Println("Executor created with ID:", newID)
		w.WriteHeader(http.StatusOK)
	}
}

func GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что метод запроса GET.
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Извлекаем переменные пути из запроса.
		vars := mux.Vars(r)

		// Проверяем наличие ключей в vars и получаем значение телефона.
		phone, ok := vars["phone"]
		if !ok || phone == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Phone not provided in request")
			return
		}

		// Создаем экземпляр newUser с данными из переменной пути.
		user := &newUser{Phone: phone}

		// Получаем токен для пользователя из API.
		token, err := api.User.GetUser(user.Phone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to get token for user: %v", err)
			return
		}

		// Устанавливаем токен как куки в ответе.
		http.SetCookie(w, &http.Cookie{
			Name:   "lovemeals", // Имя куки
			Value:  token,       // Значение куки (токен)
			Secure: true,        // Установка безопасности: куки будут передаваться только через HTTPS
			MaxAge: 3600 * 24,   // Устанавливаем срок действия куки (в секундах), здесь 24 часа
		})

		// Логгирование: выводим сообщение о выдаче токена пользователю.
		log.Printf("User with phone: %s got token", user.Phone)

		// Возвращаем статус OK.
		w.WriteHeader(http.StatusOK)
	}
}

func GetExecutor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что метод запроса GET.
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Извлекаем переменные пути из запроса.
		vars := mux.Vars(r)

		// Проверяем наличие ключей в vars и получаем значения логина и пароля.
		login, ok := vars["login"]
		if !ok || login == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Login not provided in request")
			return
		}
		password, ok := vars["password"]
		if !ok || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Password not provided in request")
			return
		}

		// Создаем экземпляр newExecutor с данными из переменных пути.
		executor := &newExecutor{
			Login:    login,
			Password: []byte(password),
		}

		// Получаем токен для исполнителя из API.
		token, err := api.Executor.GetExecutor(executor.Login, executor.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to get token for executor: %v", err)
			return
		}

		// Устанавливаем токен как куки в ответе.
		http.SetCookie(w, &http.Cookie{
			Name:   "lovemeals", // Имя куки
			Value:  token,       // Значение куки (токен)
			Secure: true,        // Установка безопасности: куки будут передаваться только через HTTPS
			MaxAge: 3600 * 24,   // Устанавливаем срок действия куки (в секундах), здесь 24 часа
		})

		// Логгирование: выводим сообщение о выдаче токена для исполнителя.
		log.Printf("Executor with login: %s got token", executor.Login)

		// Возвращаем статус OK.
		w.WriteHeader(http.StatusOK)
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что метод запроса GET.
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Удаляем куку с именем "lovemeals".
		http.SetCookie(w, &http.Cookie{
			Name:   "lovemeals", // Указываем имя куки, которую нужно удалить.
			MaxAge: -1,          // Устанавливаем отрицательное значение срока действия куки,
			// чтобы браузер удалил ее при следующем запросе.
		})

		// Возвращаем статус OK.
		w.WriteHeader(http.StatusOK)
	}
}
