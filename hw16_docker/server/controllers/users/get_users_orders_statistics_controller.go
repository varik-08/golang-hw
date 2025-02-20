package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/varik-08/golang-hw/hw16_docker/internal/entities/users"
)

func GetUsersOrdersStatistics(w http.ResponseWriter, _ *http.Request) {
	usersOrderStatistics, err := users.GetUsersOrdersStatistics()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error getting users: %s", err.Error())))
	}

	w.WriteHeader(http.StatusOK)

	resBody, err := json.Marshal(usersOrderStatistics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error marshaling users: %s", err.Error())))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
