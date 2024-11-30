package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kocannn/todos-app/middleware"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB) {

	user := r.Group("/api/v1/user")
	UserRoutes(user, db)

	todo := r.Group("/api/v1/todo")
	todo.Use(middleware.AuthMiddleware())
	TodoRoutes(todo, db)
}
