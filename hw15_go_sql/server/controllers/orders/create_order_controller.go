package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/varik-08/golang-hw/hw15_go_sql/internal/entities/orders"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error decoding Body: %s", err.Error())))
	}

	ok, validationError := validateCreateOrder(data)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationError))

		return
	}

	order := orders.OrderDTO{
		UserID:      int(data["user_id"].(float64)),
		OrderDate:   time.Now(),
		TotalAmount: data["total_amount"].(float64),
	}

	for _, p := range data["products"].([]interface{}) {
		product := p.(map[string]interface{})
		order.OrderProducts = append(order.OrderProducts, orders.OrderProductDTO{
			ProductID: int(product["product_id"].(float64)),
			Count:     int(product["count"].(float64)),
		})
	}

	id, err := orders.CreateOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error creating order: %s", err.Error())))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Order created with id: %d", id)))
}

func validateCreateOrder(data map[string]interface{}) (bool, string) {
	valid := true
	validationError := ""

	if _, ok := data["user_id"].(float64); !ok {
		validationError = "Ошибка: поле 'user_id' отсутствует или не является числом"
		valid = false
	}
	if _, ok := data["total_amount"].(float64); !ok {
		validationError = "Ошибка: поле 'total_amount' отсутствует или не является числом"
		valid = false
	}
	if _, ok := data["products"].([]interface{}); !ok {
		validationError = "Ошибка: поле 'products' отсутствует или не является числом"
		valid = false
	}
	return valid, validationError
}
