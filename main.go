package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/m-rcd/booksy/database"
	"github.com/m-rcd/booksy/handler"
	"github.com/m-rcd/booksy/models"

	"github.com/gorilla/mux"
)

func handleRequests(db database.Database) {
	h = handler.New(db)
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
	db    database.Database
	err   error
	h     handler.Handler
)

func main() {
	fmt.Println("Listening on port 10000")
	db = database.NewSQL()
	db.Open()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	handleRequests(db)
}
