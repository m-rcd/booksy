package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/m-rcd/booksy/models"
)

type SqlDB struct {
	db *sql.DB
}

var book models.Book

func NewSQL() *SqlDB {
	return &SqlDB{}
}

func (d *SqlDB) Open() error {
	db, err := sql.Open("mysql", connString)

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
	json.Unmarshal(reqBody, &book)

	_, err = stmt.Exec(book.Title, book.Author, book.Content)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (d *SqlDB) Get(id string) (models.Book, error) {
	result := d.db.QueryRow("SELECT * FROM books WHERE id = ?", id)
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
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.Content)
		if err != nil {
			return []models.Book{}, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (d *SqlDB) Delete(id string) error {
	_, err := d.db.Query("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (d *SqlDB) Update(id string, body io.ReadCloser) (models.Book, error) {
	reqBody, _ := ioutil.ReadAll(body)
	json.Unmarshal(reqBody, &book)

	_, err := d.db.Exec("UPDATE books set author=?, title=?, content=? where id=?", book.Title, book.Author, book.Content, id)
	if err != nil {
		return models.Book{}, err
	}
	fmt.Print(book)
	return book, nil
}
