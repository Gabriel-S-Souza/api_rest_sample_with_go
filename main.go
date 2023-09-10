package main

import (
	"fmt"
	"net/http"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func (p Product) toJson() string {
	return fmt.Sprintf(`{"id": %d, "name": "%s", "description": "%s", "price": %f, "quantity": %d}`, p.ID, p.Name, p.Description, p.Price, p.Quantity)
}

func main() {
	http.HandleFunc("/", getHello)
	http.HandleFunc("/products", getProducts)
	http.ListenAndServe(":8000", nil)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello World"}`))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Camiseta", Description: "Camisa azul", Price: 16.99, Quantity: 10},
		{ID: 2, Name: "Calça", Description: "Calça jeans", Price: 26.99, Quantity: 20},
		{ID: 3, Name: "Tênis", Description: "Tênis esportivo", Price: 36.99, Quantity: 30},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"products": [`))
	for i, p := range products {
		w.Write([]byte(p.toJson()))
		if i < len(products)-1 {
			w.Write([]byte(","))
		}
	}
	w.Write([]byte(`]}`))
}
