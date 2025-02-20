package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/varik-08/golang-hw/hw16_docker/internal/entities/products"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
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

	ok, validationError := validateUpdateProduct(data)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationError))

		return
	}

	product := products.ProductDTO{
		ID:    productID,
		Price: data["price"].(float64),
	}

	id, err := products.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error updating product: %s", err.Error())))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Product updated with id: %d", id)))
}

func validateUpdateProduct(data map[string]interface{}) (bool, string) {
	valid := true
	validationError := ""

	if _, ok := data["price"].(float64); !ok {
		validationError = "Ошибка: поле 'price' отсутствует или не является числом"
		valid = false
	}

	return valid, validationError
}
