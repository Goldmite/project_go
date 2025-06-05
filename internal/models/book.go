package models

import "github.com/Goldmite/project_go/internal/enums"

type Book struct {
	ISBN        string         `json:"isbn"`
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Pages       uint           `json:"pages"`
	Description string         `json:"description"`
	Publisher   string         `json:"publisher"`
	PublishDate string         `json:"publish_date"`
	Language    enums.Language `json:"language"`
}
