package handlers

import (
	"net/http"

	"github.com/Goldmite/project_shelf/internal/models"
	"github.com/Goldmite/project_shelf/internal/models/dto"
	"github.com/Goldmite/project_shelf/internal/services"
	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	groupService *services.GroupService
}

func NewGroupHandler(gs *services.GroupService) *GroupHandler {
	return &GroupHandler{groupService: gs}
}

func (groupHandler *GroupHandler) CreateGroupHandler(c *gin.Context) {
	userId := c.PostForm("id")
	groupName := c.PostForm("name")
	groupId, err := groupHandler.groupService.CreateGroup(groupName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := groupHandler.groupService.JoinGroup(userId, *groupId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, groupId)
}

func (groupHandler *GroupHandler) GetAllUserGroupsHandler(c *gin.Context) {
	userId := c.Param("id")
	groups, err := groupHandler.groupService.GetAllUserGroups(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
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

func (groupHandler *GroupHandler) GetInvitesHandler(c *gin.Context) {
	groupId := c.Param("id")
	invitedUsers, err := groupHandler.groupService.GetInvites(groupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invitedUsers)
}

func (groupHandler *GroupHandler) AcceptInvitationHandler(c *gin.Context) {
	var req dto.AcceptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := groupHandler.groupService.JoinGroup(req.UserId, req.GroupId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := groupHandler.groupService.RemoveInvite(req.UserEmail, req.GroupId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

func (groupHandler *GroupHandler) DeclineInvitationHandler(c *gin.Context) {
	var req dto.DeclineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := groupHandler.groupService.RemoveInvite(req.UserEmail, req.GroupId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}
