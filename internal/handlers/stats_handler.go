package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Goldmite/project_go/internal/models/dto"
	"github.com/Goldmite/project_go/internal/services"
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
