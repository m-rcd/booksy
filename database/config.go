package database

import (
	"fmt"
	"os"
)

var (
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	address    = "127.0.01"
	port       = "3306"
	dbName     = "books"
	connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, address, port, dbName)
)
