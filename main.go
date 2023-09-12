package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectWithDB() *sql.DB {
	godotenv.Load()
	dbPawword := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("user=postgres dbname=products_store password=%s host=localhost sslmode=disable", dbPawword)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

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
	db := connectWithDB()
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}
	products := []Product{}

	for rows.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, Product{id, name, description, price, quantity})
		defer db.Close()
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
