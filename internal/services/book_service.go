package services

import (
	"database/sql"
	"fmt"

	"github.com/Goldmite/project_go/internal/models"
)

type IBookService interface {
	CreateBook(newBook models.Book) error
	GetBookByIsbn(isbn string) (*models.Book, error)
}

type BookService struct {
	database *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{database: db}
}

func (bookService *BookService) CreateBook(b models.Book) error {
	query := "INSERT INTO books (isbn, title, author, pages, description, publisher, publish_date, language) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := bookService.database.Exec(query, b.ISBN, b.Title, b.Author, b.Pages, b.Description, b.Publisher, b.PublishDate, b.Language)
	if err != nil {
		return fmt.Errorf("failed to insert book %w", err)
	}
	return nil
}

func (bookService *BookService) GetBookByIsbn(isbn string) (*models.Book, error) {
	query := "SELECT * FROM books WHERE isbn = ?"
	row := bookService.database.QueryRow(query, isbn)

	var book models.Book
	err := row.Scan(&book.ISBN, &book.Title, &book.Author, &book.Pages, &book.Description, &book.Publisher, &book.PublishDate, &book.Language)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book with ISBN %s not found", isbn)
		}
		return nil, fmt.Errorf("failed to scan book: %w", err)
	}

	return &book, nil
}
