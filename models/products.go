package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gabriel-s-souza/api_rest_sample_with_go/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func (p Product) ToJson() string {
	return fmt.Sprintf(`{"id": %d, "name": "%s", "description": "%s", "price": %f, "quantity": %d}`, p.ID, p.Name, p.Description, p.Price, p.Quantity)
}

func productsToJson(products []Product) string {
	json := ""
	for i, p := range products {
		json += p.ToJson()
		if i < len(products)-1 {
			json += ","
		}
	}
	return json
}

func GetProducts() (string, error) {
	db := db.ConnectWithDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return "", err
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

	return fmt.Sprintf(`{"products": [%s]}`, productsToJson(products)), nil
}

func GetProductById(id int) (Product, error) {
	fmt.Println("id: ", id)
	db := db.ConnectWithDB()
	defer db.Close()
	var product Product
	sqlStatement := `SELECT * FROM products WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
	if err != nil {
		fmt.Println(err)
		if err.Error() == sql.ErrNoRows.Error() {
			return product, errors.New("Product not found: " + fmt.Sprintf("%d", id))
		} else {
			return product, err
		}
	}
	fmt.Println(product)
	return product, nil
}
