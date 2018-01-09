package config

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

func DB() *sql.DB {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db

}
