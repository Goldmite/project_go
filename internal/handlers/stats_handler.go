package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/Goldmite/project_shelf/internal/models/dto"
	"github.com/Goldmite/project_shelf/internal/services"
	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	statsService *services.StatsService
}

func NewStatsHandler(ss *services.StatsService) *StatsHandler {
	return &StatsHandler{statsService: ss}
}

func (statsHandler *StatsHandler) UpdateBookProgressHandler(c *gin.Context) {
	var request dto.BookProgressRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := statsHandler.statsService.UpdateBookProgress(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func (statsHandler *StatsHandler) GetBookProgressHandler(c *gin.Context) {
	userId := c.Query("user_id")
	isbn := c.Query("isbn")
	response, err := statsHandler.statsService.GetBookProgress(userId, isbn)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"msg": "No progress or this user doesn't have this book"})
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (statsHandler *StatsHandler) GetUserStatsHandler(c *gin.Context) {
	userId := c.Param("id")
	response, err := statsHandler.statsService.GetUserStats(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"msg": "No reading stats"})
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (statsHandler *StatsHandler) GetUserSessionsHandler(c *gin.Context) {
	userId := c.Query("user_id")
	fromDate := c.Query("from")
	if _, err := time.Parse("2006-01-02", fromDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect date format, use YYYY-MM-DD"})
		return
	}

	response, err := statsHandler.statsService.GetUserSessions(userId, fromDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
