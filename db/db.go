package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectWithDB() *sql.DB {
	godotenv.Load()
	dbPawword := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("user=postgres dbname=products_store password=%s host=localhost sslmode=disable", dbPawword)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Close(db *sql.DB) {
	db.Close()
}
