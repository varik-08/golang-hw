package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/varik-08/golang-hw/hw15_go_sql/internal/entities/products"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error decoding Body: %s", err.Error())))
	}

	ok, validationError := validateCreateProduct(data)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationError))

		return
	}

	product := products.ProductDTO{
		Name:  data["name"].(string),
		Price: data["price"].(float64),
	}

	id, err := products.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error creating product: %s", err.Error())))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Product created with id: %d", id)))
}

func validateCreateProduct(data map[string]interface{}) (bool, string) {
	valid := true
	validationError := ""

	if _, ok := data["name"].(string); !ok {
		validationError = "Ошибка: поле 'name' отсутствует или не является строкой"
		valid = false
	}
	if _, ok := data["price"].(float64); !ok {
		validationError = "Ошибка: поле 'price' отсутствует или не является числом"
		valid = false
	}

	return valid, validationError
}
