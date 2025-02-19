package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/varik-08/golang-hw/hw15_go_sql/internal/entities/users"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	usersData, err := users.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error getting users: %s", err.Error())))
	}

	w.WriteHeader(http.StatusOK)

	resBody, err := json.Marshal(usersData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error marshaling users: %s", err.Error())))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
