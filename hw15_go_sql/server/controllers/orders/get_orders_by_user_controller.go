package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/varik-08/golang-hw/hw15_go_sql/internal/entities/orders"
)

func GetOrdersByUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error decoding id: %s", err.Error())))
	}

	orderDto := orders.OrderDTO{UserID: userID}
	ordersData, err := orders.GetOrdersByUser(orderDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error getting orders: %s", err.Error())))
	}

	w.WriteHeader(http.StatusOK)

	resBody, err := json.Marshal(ordersData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error marshaling orders: %s", err.Error())))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
