package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/books", returnAllBooks)
	myRouter.HandleFunc("/book/{id}", updateBook).Methods("PATCH")
	myRouter.HandleFunc("book/", createNewBook).Methods("POST")
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

var Books []Book

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Books = []Book{
		{ID: "1", Title: "pantalaimon", Author: "Casper", Content: "hola"},
		{ID: "2", Title: "SillyCanoofy", Author: "Jasper", Content: "I am silly"},
	}
	handleRequests()
}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBooks")
	json.NewEncoder(w).Encode(Books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for _, book := range Books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for index, book := range Books {
		if book.ID == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)

	for i, b := range Books {
		if b.ID == id {
			b.Author = book.Author
			b.Content = book.Content
			b.Title = book.Title
			Books = append(Books[:i], b)
			json.NewEncoder(w).Encode(b)
		}
	}
}
