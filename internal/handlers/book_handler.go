package handlers

import (
	"fmt"
	"net/http"

	"github.com/Goldmite/project_go/internal/models/dto"
	"github.com/Goldmite/project_go/internal/services"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bs *services.BookService) *BookHandler {
	return &BookHandler{bookService: bs}
}

/*
func (bookHandler *BookHandler) CreateBookHandler(c *gin.Context) {
	isbn := c.Param("isbn")
	newBook, err := bookHandler.bookService.FetchByIsbnFromApi(isbn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bookHandler.bookService.CreateBook(*newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}*/

func (bookHandler *BookHandler) FetchByIsbnFromApiHandler(c *gin.Context) {
	isbn := c.Param("isbn")
	book, err := bookHandler.bookService.FetchByIsbnFromApi(isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
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

func (bookHandler *BookHandler) AddNewBookForUserHandler(c *gin.Context) {
	var req dto.UserBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := bookHandler.bookService.AddNewBookForUser(req.UserId, req.Isbn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

func (bookHandler *BookHandler) GetAllUserBooksHandler(c *gin.Context) {
	userId := c.Param("id")
	userBooks, err := bookHandler.bookService.GetAllUserBooks(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userBooks)
}
