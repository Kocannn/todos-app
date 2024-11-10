package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kocannn/todos-app/config"
	"github.com/kocannn/todos-app/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	db := config.ConnectionDB()

	routes.Routes(r, db)

	r.Run("localhost:8080")

}
