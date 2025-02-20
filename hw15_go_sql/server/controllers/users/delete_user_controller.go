package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/varik-08/golang-hw/hw15_go_sql/internal/entities/users"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error decoding id: %s", err.Error())))
	}

	user := users.UserDTO{
		ID: userID,
	}

	id, err := users.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error deleting user: %s", err.Error())))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User deleted with id: %d", id)))
}
