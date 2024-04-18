package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Регистрация нового пользователя
	r.HandleFunc("/users/register", NewUser()).Methods(http.MethodPost)
	// Регистрация нового исполнителя
	r.HandleFunc("/executors/register", NewExecutor()).Methods(http.MethodPost)

	// Получение токена для пользователя по номеру телефона
	r.HandleFunc("/users/login/{phone}", GetUser()).Methods(http.MethodGet)
	// Получение токена для исполнителя по логину и паролю
	r.HandleFunc("/executors/login/{login}/{password}", GetExecutor()).Methods(http.MethodGet)

	// Выход (удаление токена)
	r.HandleFunc("/get/logout", Logout()).Methods(http.MethodGet)

	return r
}
