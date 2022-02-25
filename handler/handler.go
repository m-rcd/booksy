package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/m-rcd/booksy/database"
	"github.com/m-rcd/booksy/models"
)

type Handler struct {
	db database.Database
}

func New(db database.Database) Handler {
	return Handler{db: db}
}

func (h *Handler) ReturnSingleBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	w.Header().Set("Content-Type", "application/json")

	book, err := h.db.Get(id)

	if err != nil {
		response := failedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	} else {
		json.NewEncoder(w).Encode(book)
	}

}

func (h *Handler) CreateNewBook(w http.ResponseWriter, r *http.Request) {
	var response models.JsonBookResponse

	newBook, err := h.db.Create(r.Body)
	if err != nil {
		response = failedResponse(err.Error())
	} else {
		response = successResponse([]models.Book{newBook}, "The book was successfully created")
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var response models.JsonBookResponse

	id := mux.Vars(r)["id"]
	book, err := h.db.Update(id, r.Body)
	if err != nil {
		response = failedResponse(err.Error())
	} else {
		response = successResponse([]models.Book{book}, "The book was successfully updated")
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ReturnAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := h.db.List()
	if err != nil {
		response := failedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	} else {
		json.NewEncoder(w).Encode(books)
	}

}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.db.Delete(id)
	var response models.JsonBookResponse

	if err != nil {
		response = failedResponse(err.Error())
	} else {
		response = successResponse([]models.Book{}, "The book was successfully deleted")
	}

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func failedResponse(message string) models.JsonBookResponse {
	return models.JsonBookResponse{Type: "failed", StatusCode: 500, Data: []models.Book{}, Message: message}
}

func successResponse(data []models.Book, message string) models.JsonBookResponse {
	return models.JsonBookResponse{Type: "success", StatusCode: 200, Data: data, Message: message}
}
