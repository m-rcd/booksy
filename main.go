package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/m-rcd/go-rest-api/handler"
	"github.com/m-rcd/go-rest-api/models"

	"github.com/gorilla/mux"
)

func handleRequests(db *sql.DB) {
	h = h.New(db)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", h.HomePage)
	myRouter.HandleFunc("/books", h.ReturnAllBooks)
	myRouter.HandleFunc("/book/{id}", h.UpdateBook).Methods("PATCH")
	myRouter.HandleFunc("/book", h.CreateNewBook).Methods("POST")
	myRouter.HandleFunc("/book/{id}", h.DeleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", h.ReturnSingleBook)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

var (
	Books []models.Book
	db    *sql.DB
	err   error
	h     handler.Handler
)

func main() {
	fmt.Println("Listening on port 10000")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connection := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/books", dbUsername, dbPassword)
	db, err = sql.Open("mysql", connection)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	handleRequests(db)
}
