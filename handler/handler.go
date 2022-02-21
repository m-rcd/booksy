package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/m-rcd/go-rest-api/models"
)

type Handler struct {
	db *sql.DB
}

var (
	response models.JsonBookResponse
	Books    []models.Book
)

func (h *Handler) New(db *sql.DB) Handler {
	return Handler{db: db}
}

func (h *Handler) ReturnSingleBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var book models.Book
	w.Header().Set("Content-Type", "application/json")
	result := h.db.QueryRow("SELECT * FROM books WHERE id = ?", id)

	err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Content)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(book)
}

func (h *Handler) CreateNewBook(w http.ResponseWriter, r *http.Request) {
	stmt, err := h.db.Prepare("INSERT INTO books(title, author, content) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newBook models.Book
	json.Unmarshal(reqBody, &newBook)

	_, err = stmt.Exec(newBook.Title, newBook.Author, newBook.Content)
	if err != nil {
		panic(err.Error())
	}
	response = models.JsonBookResponse{Type: "sucess", Data: []models.Book{newBook}, Message: "The book was successfully created"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book models.Book
	json.Unmarshal(reqBody, &book)

	h.db.Exec("UPDATE books set author=?, title=?, content=? where id=?", book.Title, book.Author, book.Content, id)
	response = models.JsonBookResponse{Type: "success", Data: []models.Book{book}, Message: "The book was successfully updated"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ReturnAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := h.db.Query("SELECT * from books")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var book models.Book
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Content)
		if err != nil {
			panic(err.Error())
		}
		Books = append(Books, book)
	}
	json.NewEncoder(w).Encode(Books)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := h.db.Query("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	response = models.JsonBookResponse{Type: "success", Message: "The book has been deleted successfully!"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
