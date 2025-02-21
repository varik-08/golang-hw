package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/varik-08/golang-hw/hw16_docker/internal/entities/users"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error decoding id: %s", err.Error())))
	}

	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error decoding Body: %s", err.Error())))
	}

	ok, validationError := validateUpdateUser(data)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationError))

		return
	}

	user := users.UserDTO{
		ID:   userID,
		Name: data["name"].(string),
	}

	id, err := users.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error updating user: %s", err.Error())))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User updated with id: %d", id)))
}

func validateUpdateUser(data map[string]interface{}) (bool, string) {
	valid := true
	validationError := ""

	if _, ok := data["name"].(string); !ok {
		validationError = "Ошибка: поле 'name' отсутствует или не является строкой"
		valid = false
	}

	return valid, validationError
}
