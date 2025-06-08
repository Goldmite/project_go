package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
	authorsJSON, err := json.Marshal(b.Authors)
	if err != nil {
		return err
	}
	query := "INSERT INTO books (isbn, title, authors, pages, description, publisher, publish_date, language) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = bookService.database.Exec(query, b.ISBN, b.Title, string(authorsJSON), b.Pages, b.Description, b.Publisher, b.PublishDate, b.Language)
	if err != nil {
		return fmt.Errorf("failed to insert book %w", err)
	}
	return nil
}
func (bookService *BookService) FetchByIsbnFromApi(isbn string) (*models.Book, error) {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s", isbn, apiKey)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result models.BooksResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	if len(result.Items) == 0 {
		return nil, nil
	}
	return &result.Items[0].Book, nil
}

func (bookService *BookService) GetBookByIsbn(isbn string) (*models.Book, error) {
	query := "SELECT * FROM books WHERE isbn = ?"
	row := bookService.database.QueryRow(query, isbn)

	var book models.Book
	err := row.Scan(&book.ISBN, &book.Title, &book.Authors, &book.Pages, &book.Description, &book.Publisher, &book.PublishDate, &book.Language)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book with ISBN %s not found", isbn)
		}
		return nil, fmt.Errorf("failed to scan book: %w", err)
	}

	return &book, nil
}

func (bookService *BookService) AddNewBookForUser(userId, isbn string) error {
	// Book does not exist -> add it
	if _, err := bookService.GetBookByIsbn(isbn); err != nil {
		newBook, err := bookService.FetchByIsbnFromApi(isbn)
		if err != nil {
			return err
		}
		newBook.ISBN = isbn

		if err := bookService.CreateBook(*newBook); err != nil {
			return err
		}
	}

	query := "INSERT INTO views (user_id, book_id) VALUES (?, ?)"
	_, err := bookService.database.Exec(query, userId, isbn)
	if err != nil {
		return fmt.Errorf("failed to insert view %w", err)
	}

	return nil
}

func (bookService *BookService) GetAllUserBooks(userId string) ([]models.Book, error) {
	query := "SELECT isbn, title, authors, pages, description, publisher, publish_date, language FROM books JOIN views ON isbn = book_id WHERE user_id = ?"
	rows, err := bookService.database.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userBooks []models.Book
	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ISBN, &b.Title, &b.Authors, &b.Pages, &b.Description, &b.Publisher, &b.PublishDate, &b.Language)
		if err != nil {
			return nil, err
		}
		userBooks = append(userBooks, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userBooks, nil
}
