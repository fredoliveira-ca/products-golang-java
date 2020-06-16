package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	defaultHost = "localhost"
)

//ConnectDataBase is a way to open a connection with the database.
// If the environment variable is not available, it must assume the default value.
func ConnectDataBase() *sql.DB {
	host := defaultHost
	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}

	connection := fmt.Sprintf(
		"host=%s port=%s user=%s  password=%s dbname=%s sslmode=disable",
		host,
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	log.Println("Connecting database", connection)
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
