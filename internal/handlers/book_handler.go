package handlers

import (
	"fmt"
	"net/http"

	"github.com/Goldmite/project_go/internal/models"
	"github.com/Goldmite/project_go/internal/services"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bs *services.BookService) *BookHandler {
	return &BookHandler{bookService: bs}
}

func (bookHandler *BookHandler) CreateBookHandler(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := bookHandler.bookService.CreateBook(newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func (bookHandler *BookHandler) GetBookByIsbnHandler(c *gin.Context) {
	isbn := c.Param("isbn")
	book, err := bookHandler.bookService.GetBookByIsbn(isbn)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("book with ISBN %s not found", isbn)})
		return
	}

	c.JSON(http.StatusOK, book)
}
