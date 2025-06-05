package services_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Goldmite/project_go/internal/enums"
	"github.com/Goldmite/project_go/internal/models"
	"github.com/Goldmite/project_go/internal/services"
	"github.com/stretchr/testify/assert"
)

var mockBook = models.Book{
	ISBN:        "9786094798269",
	Title:       "Procesas ir novelės",
	Author:      "Franz Kafka",
	Pages:       391,
	Description: "Testing lalalalalala",
	Publisher:   "baltos lankos",
	PublishDate: "2023-01-01",
	Language:    enums.Lithuanian,
}

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when connecting db", err)
	}
	defer db.Close()
	bookService := services.NewBookService(db)

	tests := []struct {
		name     string
		newBook  models.Book
		mockFunc func()
		expected error
	}{
		{
			name:    "Created successfully",
			newBook: mockBook,
			mockFunc: func() {
				mock.ExpectExec("INSERT INTO books \\(isbn, title, author, pages, description, publisher, publish_date, language\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").
					WithArgs(mockBook.ISBN, mockBook.Title, mockBook.Author, mockBook.Pages, mockBook.Description, mockBook.Publisher, mockBook.PublishDate, mockBook.Language).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expected: nil,
		},
		{
			name:    "Failed creation due to DB error",
			newBook: mockBook,
			mockFunc: func() {
				mock.ExpectExec("INSERT INTO books \\(isbn, title, author, pages, description, publisher, publish_date, language\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").
					WithArgs(mockBook.ISBN, mockBook.Title, mockBook.Author, mockBook.Pages, mockBook.Description, mockBook.Publisher, mockBook.PublishDate, mockBook.Language).
					WillReturnError(errors.New("database error"))
			},
			expected: errors.New("failed to insert book"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			err := bookService.CreateBook(tc.newBook)
			if tc.expected != nil {
				assert.ErrorContains(t, err, tc.expected.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

func TestGetBookByIsbn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error '%s' when connecting db", err)
	}
	defer db.Close()
	bookService := services.NewBookService(db)

	tests := []struct {
		name     string
		isbn     string
		mockFunc func()
		res      *models.Book
		err      error
	}{
		{
			name: "Found book",
			isbn: mockBook.ISBN,
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{"isbn", "title", "author", "pages", "description", "publisher", "publish_date", "language"}).
					AddRow(mockBook.ISBN, mockBook.Title, mockBook.Author, mockBook.Pages, mockBook.Description, mockBook.Publisher, mockBook.PublishDate, mockBook.Language)
				mock.ExpectQuery("SELECT \\* FROM books WHERE isbn = \\?").
					WithArgs(mockBook.ISBN).
					WillReturnRows(rows)
			},
			res: &mockBook,
			err: nil,
		},
		{
			name: "Book not found",
			isbn: "12121212120",
			mockFunc: func() {
				mock.ExpectQuery("SELECT \\* FROM books WHERE isbn = \\?").
					WithArgs("12121212120").
					WillReturnError(sql.ErrNoRows)
			},
			res: nil,
			err: errors.New("book with ISBN 12121212120 not found"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			book, err := bookService.GetBookByIsbn(tc.isbn)
			if tc.err != nil {
				assert.ErrorContains(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.res, book)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
