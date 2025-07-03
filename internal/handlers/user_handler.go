package handlers

import (
	"net/http"
	"strings"

	"github.com/Goldmite/project_go/internal/models"
	"github.com/Goldmite/project_go/internal/models/dto"
	"github.com/Goldmite/project_go/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (userHandler *UserHandler) CreateUserHandler(c *gin.Context) {
	var request models.CreateUserRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := userHandler.userService.CreateUser(request)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists: " + err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (userHandler *UserHandler) GetUserHandler(c *gin.Context) {
	var res *dto.GetUserResponse
	var err error
	email := c.Param("email")
	if email != "" {
		res, err = userHandler.userService.GetUserByIdOrEmail("", email)
	} else {
		email := c.PostForm("email")
		password := c.PostForm("password")
		if email == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
		}

		res, err = userHandler.userService.GetUserByLogin(email, password)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) GetUserInvitesHandler(c *gin.Context) {
	userId := c.Param("id")
	res, err := userHandler.userService.GetUserInvites(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (userHandler *UserHandler) GetGroupMembersHandler(c *gin.Context) {
	groupId := c.Param("id")
	res, err := userHandler.userService.GetGroupMembers(groupId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
