package models

import "github.com/Goldmite/project_go/internal/enums"

type Book struct {
	ISBN        string         `json:"isbn"`
	Title       string         `json:"title"`
	Authors     []string       `json:"authors"`
	Pages       uint           `json:"pageCount"`
	Description string         `json:"description"`
	Publisher   string         `json:"publisher"`
	PublishDate string         `json:"publishedDate"`
	Language    enums.Language `json:"language"`
}

type BookItem struct {
	Book Book `json:"volumeInfo"`
}

type BooksResponse struct {
	Items []BookItem `json:"items"`
}
