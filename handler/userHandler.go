package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/todos-app/domain"
)

type UserHandler struct {
	UserService domain.UserService
}

// CreateUser implements domain.UserHandler.
func (u *UserHandler) CreateUser(c *gin.Context) {
	panic("unimplemented")
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
