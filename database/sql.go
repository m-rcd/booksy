package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/m-rcd/go-rest-api/models"
)

type SqlDB struct {
	db *sql.DB
}

func NewSQL() *SqlDB {
	return &SqlDB{}
}

func (d *SqlDB) Open() error {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connection := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/books", dbUsername, dbPassword)
	db, err := sql.Open("mysql", connection)

	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *SqlDB) Close() error {
	return d.db.Close()
}

func (d *SqlDB) Create(body io.ReadCloser) (models.Book, error) {
	stmt, err := d.db.Prepare("INSERT INTO books(title, author, content) VALUES(?, ?, ?)")
	if err != nil {
		return models.Book{}, err
	}

	reqBody, _ := ioutil.ReadAll(body)
	var newBook models.Book
	json.Unmarshal(reqBody, &newBook)

	_, err = stmt.Exec(newBook.Title, newBook.Author, newBook.Content)
	if err != nil {
		return models.Book{}, err
	}

	return newBook, nil
}

func (d *SqlDB) Get(id string) (models.Book, error) {
	result := d.db.QueryRow("SELECT * FROM books WHERE id = ?", id)
	var book models.Book
	result.Scan(&book.ID, &book.Title, &book.Author, &book.Content)

	return book, nil
}

func (d *SqlDB) List() ([]models.Book, error) {
	result, err := d.db.Query("SELECT * FROM books")
	if err != nil {
		return []models.Book{}, err
	}
	defer result.Close()

	var books []models.Book

	for result.Next() {
		var book models.Book
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Content)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}
	return books, nil
}

func (d *SqlDB) Delete(id string) error {
	_, err := d.db.Query("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func (d *SqlDB) Update(id string, body io.ReadCloser) (models.Book, error) {
	reqBody, _ := ioutil.ReadAll(body)
	var book models.Book
	json.Unmarshal(reqBody, &book)

	_, err := d.db.Exec("UPDATE books set author=?, title=?, content=? where id=?", book.Title, book.Author, book.Content, id)
	if err != nil {
		return models.Book{}, err
	}
	fmt.Print(book)
	return book, nil
}
