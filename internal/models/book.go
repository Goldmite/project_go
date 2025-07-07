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
	Cover       Image          `json:"imageLinks"`
}

type Image struct {
	Url string `json:"thumbnail"`
}

type BookItem struct {
	Book Book `json:"volumeInfo"`
}

type BookQueryResponse struct {
	Items []BookItem `json:"items"`
}
