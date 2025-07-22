package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Goldmite/project_shelf/internal/models"
	"github.com/Goldmite/project_shelf/internal/models/dto"
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
	query := "INSERT INTO books (isbn, title, authors, pages, description, publisher, publish_date, language, cover_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = bookService.database.Exec(query, b.ISBN, b.Title, string(authorsJSON), b.Pages, b.Description, b.Publisher, b.PublishDate, b.Language, b.Cover.Url)
	if err != nil {
		return fmt.Errorf("failed to insert book: %w", err)
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

	var result models.BookQueryResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("book not found with ISBN: %s", isbn)
	}
	return &result.Items[0].Book, nil
}

func (bookService *BookService) GetBookByIsbn(isbn string) (*dto.BookInfoResponse, error) {
	query := "SELECT * FROM books WHERE isbn = ?"
	row := bookService.database.QueryRow(query, isbn)

	var book dto.BookInfoResponse
	var authorsJson string
	err := row.Scan(&book.ISBN, &book.Title, &authorsJson, &book.Pages, &book.Description, &book.Publisher, &book.PublishDate, &book.Language, &book.Cover)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(authorsJson), &book.Authors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse authors JSON: %w", err)
	}
	return &book, nil
}

func (bookService *BookService) GetRecentlyReadBook(userId string) (*dto.RecentlyReadResponse, error) {
	query := "SELECT isbn, title, authors, cover_url, session_created_at, session_updated_at FROM books " +
		"JOIN reading ON isbn = book_id " +
		"WHERE user_id = ? " +
		"ORDER BY session_updated_at DESC " +
		"LIMIT 1"
	row := bookService.database.QueryRow(query, userId)

	var (
		recent      dto.RecentlyReadResponse
		authorsJson string
	)
	err := row.Scan(&recent.ISBN, &recent.Title, &authorsJson, &recent.Cover, &recent.StartDate, &recent.LastReadDate)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(authorsJson), &recent.Authors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse authors JSON: %w", err)
	}

	return &recent, nil
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

	query := "INSERT INTO reading (user_id, book_id) VALUES (?, ?)"
	_, err := bookService.database.Exec(query, userId, isbn)
	if err != nil {
		return fmt.Errorf("failed to insert view %w", err)
	}

	return nil
}

func (bookService *BookService) GetAllUserBooks(userId string) ([]dto.BookResponse, error) {
	query := "SELECT user_id, isbn, title, authors, cover_url FROM books JOIN reading ON isbn = book_id WHERE user_id = ?"
	rows, err := bookService.database.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userBooks, err := GetAllBooksHelper(rows)
	if err != nil {
		return nil, err
	}

	return userBooks, nil
}

func (bookService *BookService) GetAllGroupBooks(groupId string) ([]dto.BookResponse, error) {
	query := "SELECT GROUP_CONCAT(m.user_id), isbn, title, authors, cover_url FROM books b " +
		"JOIN reading r ON b.isbn = r.book_id " +
		"JOIN members m ON r.user_id = m.user_id " +
		"WHERE m.group_id = ? GROUP BY isbn"
	rows, err := bookService.database.Query(query, groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groupBooks, err := GetAllBooksHelper(rows)
	if err != nil {
		return nil, err
	}

	return groupBooks, nil
}

func GetAllBooksHelper(rows *sql.Rows) ([]dto.BookResponse, error) {
	var books []dto.BookResponse
	for rows.Next() {
		var b dto.BookResponse
		var groupedOwners string
		var authorsJsons string
		err := rows.Scan(&groupedOwners, &b.ISBN, &b.Title, &authorsJsons, &b.Cover)
		if err != nil {
			return nil, err
		}
		b.OwnedBy = strings.Split(groupedOwners, ",")

		err = json.Unmarshal([]byte(authorsJsons), &b.Authors)
		if err != nil {
			return nil, fmt.Errorf("failed to parse authors JSON: %w", err)
		}
		books = append(books, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
