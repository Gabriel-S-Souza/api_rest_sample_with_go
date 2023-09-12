package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gabriel-s-souza/api_rest_sample_with_go/models"
	"github.com/gorilla/mux"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello World"}`))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the products"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(products))
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid product ID"}`))
		return
	}

	product, err := models.GetProductById(id)

	if err != nil {
		if strings.Contains(err.Error(), "product not found") {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(product.ToJson()))
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid body"}`))
		return
	}

	var product models.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Error unmarshalling the body"}`))
		return
	}

	productResp, err := product.CreateProduct(product.Name, product.Description, product.Price, product.Quantity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(productResp.ToJson()))
}
