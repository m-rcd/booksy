package responses

import "github.com/m-rcd/booksy/models"

type Response interface {
	FailedResponse(message string) JsonBookResponse
	SuccessResponse(data []models.Book, message string) JsonBookResponse
}
