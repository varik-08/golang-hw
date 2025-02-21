package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/varik-08/golang-hw/hw16_docker/internal/entities/users"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error decoding Body: %s", err.Error())))
	}

	ok, validationError := validateCreateUser(data)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationError))

		return
	}

	user := users.UserDTO{
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}

	id, err := users.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error creating user: %s", err.Error())))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created with id: %d", id)))
}

func validateCreateUser(data map[string]interface{}) (bool, string) {
	valid := true
	validationError := ""

	if _, ok := data["name"].(string); !ok {
		validationError = "Ошибка: поле 'name' отсутствует или не является строкой"
		valid = false
	}
	if _, ok := data["email"].(string); !ok {
		validationError = "Ошибка: поле 'email' отсутствует или не является строкой"
		valid = false
	}
	if _, ok := data["password"].(string); !ok {
		validationError = "Ошибка: поле 'password' отсутствует или не является строкой"
		valid = false
	}

	return valid, validationError
}
