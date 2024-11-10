package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kocannn/todos-app/config"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	config.ConnectionDB()

	r.Run("localhost:8080")

}
