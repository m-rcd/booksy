package models

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type JsonBookResponse struct {
	Type    string `json:"type"`
	Data    []Book `json:"data"`
	Message string `json:"message"`
}
