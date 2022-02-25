package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/m-rcd/booksy/database/databasefakes"
	"github.com/m-rcd/booksy/handler"
	"github.com/m-rcd/booksy/models"
	"github.com/m-rcd/booksy/pkg/responses"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var jsonResponse responses.JsonBookResponse

var _ = Describe("Handler", func() {
	Context("POST request", func() {
		It("handles CreateNewBook request", func() {
			fake_db := new(databasefakes.FakeDatabase)

			data := bytes.NewBuffer([]byte(`{"title":"Northern lights", "author": "Philip Pullman", "content":"iexist"}`))
			req, err := http.NewRequest("POST", "http://localhost:10000/book", data)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)

			book := models.Book{Title: "Northern lights", Author: "Philip Pullman", Content: "iexist"}
			fake_db.CreateReturns(book, nil)
			h.CreateNewBook(rr, req)
			Expect(fake_db.CreateCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("success"))
			Expect(jsonResponse.StatusCode).To(Equal(200))
			Expect(jsonResponse.Data[0].Title).To(Equal("Northern lights"))
			Expect(jsonResponse.Data[0].Author).To(Equal("Philip Pullman"))
			Expect(jsonResponse.Data[0].Content).To(Equal("iexist"))
			Expect(jsonResponse.Message).To(Equal("The book was successfully created"))
		})

		It("returns error when CreateNewBoook request fails", func() {
			fake_db := new(databasefakes.FakeDatabase)

			data := bytes.NewBuffer([]byte(`{"key": 9, "author": "writer", "content":"iexist"}`))
			req, err := http.NewRequest("POST", "http://localhost:10000/book", data)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)

			book := models.Book{}
			fake_db.CreateReturns(book, errors.New("Creating book failed"))
			h.CreateNewBook(rr, req)
			Expect(fake_db.CreateCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("failed"))
			Expect(jsonResponse.StatusCode).To(Equal(500))
			Expect(len(jsonResponse.Data)).To(Equal(0))
			Expect(jsonResponse.Message).To(Equal("Creating book failed"))
		})
	})

	Context("GET request", func() {
		It("handles ReturnSingleBook", func() {
			fake_db := new(databasefakes.FakeDatabase)

			req, err := http.NewRequest("GET", "http://localhost:10000/book/1", nil)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)
			book := models.Book{ID: "1", Title: "The subtle knife", Author: "Philip Pullman", Content: "iexist"}

			fake_db.GetReturns(book, nil)
			h.ReturnSingleBook(rr, req)
			Expect(fake_db.GetCallCount()).To(Equal(1))
			var getBook models.Book
			json.Unmarshal(rr.Body.Bytes(), &getBook)
			Expect(getBook.ID).To(Equal(book.ID))
			Expect(getBook.Content).To(Equal(book.Content))
			Expect(getBook.Title).To(Equal(book.Title))
			Expect(getBook.Author).To(Equal(book.Author))
		})

		It("returns error when ReturnSingleBook fails", func() {
			fake_db := new(databasefakes.FakeDatabase)

			req, err := http.NewRequest("GET", "http://localhost:10000/book/0", nil)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)
			book := models.Book{}

			fake_db.GetReturns(book, errors.New("Getting book failed"))
			h.ReturnSingleBook(rr, req)
			Expect(fake_db.GetCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("failed"))
			Expect(jsonResponse.StatusCode).To(Equal(500))
			Expect(len(jsonResponse.Data)).To(Equal(0))
			Expect(jsonResponse.Message).To(Equal("Getting book failed"))
		})
	})

	Context("LIST request", func() {
		It("handles ReturnAllBooks", func() {
			fake_db := new(databasefakes.FakeDatabase)

			req, err := http.NewRequest("GET", "http://localhost:10000/books", nil)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)
			book1 := models.Book{ID: "1", Title: "Book1", Author: "Author1", Content: "iexist1"}
			book2 := models.Book{ID: "2", Title: "Book2", Author: "Author2", Content: "iexist2"}
			books := []models.Book{book1, book2}
			fake_db.ListReturns(books, nil)
			h.ReturnAllBooks(rr, req)
			Expect(fake_db.ListCallCount()).To(Equal(1))

			var list []models.Book
			json.Unmarshal(rr.Body.Bytes(), &list)
			Expect(list[0].ID).To(Equal(book1.ID))
			Expect(list[0].Content).To(Equal(book1.Content))
			Expect(list[0].Title).To(Equal(book1.Title))
			Expect(list[0].Author).To(Equal(book1.Author))
			Expect(len(list)).To(Equal(2))
		})

		It("returns error when ReturnAllBooks fails", func() {
			fake_db := new(databasefakes.FakeDatabase)

			req, err := http.NewRequest("GET", "http://localhost:10000/books", nil)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)

			fake_db.ListReturns([]models.Book{}, errors.New("Getting books failed"))
			h.ReturnAllBooks(rr, req)
			Expect(fake_db.ListCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("failed"))
			Expect(jsonResponse.StatusCode).To(Equal(500))
			Expect(len(jsonResponse.Data)).To(Equal(0))
			Expect(jsonResponse.Message).To(Equal("Getting books failed"))
		})
	})

	Context("DELETE request", func() {
		It("handles DeleteBook request", func() {
			fake_db := new(databasefakes.FakeDatabase)

			req, err := http.NewRequest("DELETE", "http://localhost:10000/book/1", nil)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)

			fake_db.DeleteReturns(nil)
			h.DeleteBook(rr, req)
			Expect(fake_db.DeleteCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("success"))
			Expect(jsonResponse.StatusCode).To(Equal(200))
			Expect(len(jsonResponse.Data)).To(Equal(0))
			Expect(jsonResponse.Message).To(Equal("The book was successfully deleted"))
		})

		It("returns error when DeleteBook request fails", func() {
			fake_db := new(databasefakes.FakeDatabase)

			req, err := http.NewRequest("DELETE", "http://localhost:10000/book/1", nil)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)

			fake_db.DeleteReturns(errors.New("Error while deleting"))
			h.DeleteBook(rr, req)
			Expect(fake_db.DeleteCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("failed"))
			Expect(jsonResponse.StatusCode).To(Equal(500))
			Expect(len(jsonResponse.Data)).To(Equal(0))
			Expect(jsonResponse.Message).To(Equal("Error while deleting"))
		})
	})

	Context("PATCH request", func() {
		It("handles UpdateBook request", func() {
			fake_db := new(databasefakes.FakeDatabase)

			data := bytes.NewBuffer([]byte(`{"title":"hello", "author": "writer", "content":"iexist"}`))
			req, err := http.NewRequest("PATCH", "http://localhost:10000/book/1", data)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)

			book := models.Book{Title: "hello", Author: "writer", Content: "iexist"}
			fake_db.UpdateReturns(book, nil)
			h.UpdateBook(rr, req)
			Expect(fake_db.UpdateCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("success"))
			Expect(jsonResponse.StatusCode).To(Equal(200))
			Expect(jsonResponse.Data[0].Title).To(Equal("hello"))
			Expect(jsonResponse.Data[0].Author).To(Equal("writer"))
			Expect(jsonResponse.Data[0].Content).To(Equal("iexist"))
			Expect(jsonResponse.Message).To(Equal("The book was successfully updated"))
		})

		It("returns error when UpdateBook request fails", func() {
			fake_db := new(databasefakes.FakeDatabase)

			data := bytes.NewBuffer([]byte(`{"key": 9, "author": "writer", "content":"iexist"}`))
			req, err := http.NewRequest("POST", "http://localhost:10000/book/1", data)
			Expect(err).NotTo(HaveOccurred())
			rr := httptest.NewRecorder()
			h := handler.New(fake_db)

			book := models.Book{}
			fake_db.UpdateReturns(book, errors.New("Updating book failed"))
			h.UpdateBook(rr, req)
			Expect(fake_db.UpdateCallCount()).To(Equal(1))

			json.Unmarshal(rr.Body.Bytes(), &jsonResponse)
			Expect(jsonResponse.Type).To(Equal("failed"))
			Expect(jsonResponse.StatusCode).To(Equal(500))
			Expect(len(jsonResponse.Data)).To(Equal(0))
			Expect(jsonResponse.Message).To(Equal("Updating book failed"))
		})

	})
})
