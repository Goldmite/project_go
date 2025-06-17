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
	if request.Name == "" {
		request.Name = strings.Split(request.Email, "@")[0]
	}
	id, err := userHandler.userService.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (userHandler *UserHandler) GetUserHandler(c *gin.Context) {
	var res *dto.GetUserResponse
	var err error
	id := c.Param("id")
	if id != "" {
		res, err = userHandler.userService.GetUserById(id)
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

func (userHandler *UserHandler) AcceptInvitationHandler(c *gin.Context) {
	var req dto.AcceptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userHandler.userService.JoinGroup(req.UserId, req.GroupId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := userHandler.userService.RemoveInvite(req.Token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

func (userHandler *UserHandler) DeclineInvitationHandler(c *gin.Context) {
	token := c.Param("token")
	if err := userHandler.userService.RemoveInvite(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
