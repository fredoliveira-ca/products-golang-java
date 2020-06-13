package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "productdb"
)

//ConnectDataBase is a way to open a connection with the database
func ConnectDataBase() *sql.DB {
	connection := fmt.Sprintf("host=%s port=%d user=%s  password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
