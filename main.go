package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/m-rcd/booksy/pkg/database"
	"github.com/m-rcd/booksy/pkg/handler"
	"github.com/m-rcd/booksy/pkg/models"

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
	myRouter.HandleFunc("/book/{id}", h.ReturnSingleBook).Methods("GET")

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

	username := os.Getenv("DB_USERNAME")
	if username == "" {
		fmt.Println("DB_USERNAME must be set")
		os.Exit(1)
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		fmt.Println("DB_PASSWORD must be set")
		os.Exit(1)
	}

	db = database.NewSQL(username, password, database.Address, database.Port)
	err = db.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	handleRequests(db)

}
