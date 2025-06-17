package handlers

import (
	"net/http"

	"github.com/Goldmite/project_go/internal/models"
	"github.com/Goldmite/project_go/internal/models/dto"
	"github.com/Goldmite/project_go/internal/services"
	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	groupService *services.GroupService
}

func NewGroupHandler(gs *services.GroupService) *GroupHandler {
	return &GroupHandler{groupService: gs}
}

func (groupHandler *GroupHandler) CreateGroupHandler(c *gin.Context) {
	groupName := c.Param("name")
	id, err := groupHandler.groupService.CreateGroup(groupName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (groupHandler *GroupHandler) SendInvitesHandler(c *gin.Context) {
	var req dto.InviteRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for nr := range req.EmailTo {
		invite := models.NewInvitationFromRequest(req, nr)
		err := groupHandler.groupService.CreateInvite(*invite)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, req)
}
