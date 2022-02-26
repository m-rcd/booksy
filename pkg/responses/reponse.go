package responses

import "github.com/m-rcd/booksy/pkg/models"

type Response interface {
	Failure(message string) JsonBookResponse
	Success(data []models.Book, message string) JsonBookResponse
}
