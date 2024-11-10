package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kocannn/todos-app/domain"
)

type UserHandler struct {
	UserService domain.UserService
}

// CreateUser implements domain.UserHandler.
func (u *UserHandler) CreateUser(c *gin.Context) {
	var user *domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createUser, err := u.UserService.CreateUser(*user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": createUser})

}

// Login implements domain.UserHandler.
func (u *UserHandler) Login(c *gin.Context) {
	panic("unimplemented")
}

func NewUserHandler(userService domain.UserService) domain.UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
