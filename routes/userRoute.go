package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/todos-app/handler"
	"github.com/kocannn/todos-app/repository"
	"github.com/kocannn/todos-app/service"
	"gorm.io/gorm"
)

func UserRoutes(group *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	group.POST("/signup", userHandler.CreateUser)
}
