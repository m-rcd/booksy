package responses

import "github.com/m-rcd/booksy/pkg/models"

type JsonBookResponse struct {
	Type       string        `json:"type"`
	StatusCode int           `json:status_code`
	Data       []models.Book `json:"data"`
	Message    string        `json:"message"`
}

type BookResponse struct {
	Response JsonBookResponse
}

func NewBookResponse() *BookResponse {
	return &BookResponse{}
}

func (b *BookResponse) Failure(message string) JsonBookResponse {
	b.Response = JsonBookResponse{Type: "failed", StatusCode: 500, Data: []models.Book{}, Message: message}
	return b.Response
}

func (b *BookResponse) Success(data []models.Book, message string) JsonBookResponse {
	b.Response = JsonBookResponse{Type: "success", StatusCode: 200, Data: data, Message: message}
	return b.Response
}
