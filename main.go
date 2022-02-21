package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", returnAllBooks)
	myRouter.HandleFunc("/book/{id}", updateBook).Methods("PATCH")
	myRouter.HandleFunc("/book", createNewBook).Methods("POST")
	myRouter.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", returnSingleBook)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Book `json:"data"`
	Message string `json:"message"`
}

var (
	Books    []Book
	db       *sql.DB
	err      error
	response JsonResponse
)

func main() {
	fmt.Println("Listening on port 10000")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connection := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/books", dbUsername, dbPassword)
	db, err = sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	handleRequests()
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT * from books")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var book Book
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Content)
		if err != nil {
			panic(err.Error())
		}
		Books = append(Books, book)
	}
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var book Book
	w.Header().Set("Content-Type", "application/json")
	result := db.QueryRow("SELECT * FROM books WHERE id = ?", id)

	err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Content)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(book)
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO books(title, author, content) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newBook Book
	json.Unmarshal(reqBody, &newBook)

	_, err = stmt.Exec(newBook.Title, newBook.Author, newBook.Content)
	if err != nil {
		panic(err.Error())
	}
	response = JsonResponse{Type: "sucess", Data: []Book{newBook}, Message: "The book was successfully created"}
	json.NewEncoder(w).Encode(response)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.Query("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	response = JsonResponse{Type: "success", Message: "The book has been deleted successfully!"}
	json.NewEncoder(w).Encode(response)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)

	db.Exec("UPDATE books set author=?, title=?, content=? where id=?", book.Title, book.Author, book.Content, id)
	response = JsonResponse{Type: "success", Data: []Book{book}, Message: "The book was successfully updated"}
	json.NewEncoder(w).Encode(response)
}
