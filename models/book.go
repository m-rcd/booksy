package models

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type JsonBookResponse struct {
	Type       string `json:"type"`
	StatusCode int    `json:status_code`
	Data       []Book `json:"data"`
	Message    string `json:"message"`
}
