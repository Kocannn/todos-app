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
	createUser, err := u.UserService.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": createUser})

}

// Login implements domain.UserHandler.
func (u *UserHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindBodyWithJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := u.UserService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login success"})
}

func NewUserHandler(userService domain.UserService) domain.UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
