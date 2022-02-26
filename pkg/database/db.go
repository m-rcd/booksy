package database

import (
	"io"

	"github.com/m-rcd/booksy/pkg/models"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Database
type Database interface {
	Open() error
	Close() error
	Create(body io.ReadCloser) (models.Book, error)
	Get(id string) (models.Book, error)
	List() ([]models.Book, error)
	Delete(id string) error
	Update(id string, body io.ReadCloser) (models.Book, error)
}
