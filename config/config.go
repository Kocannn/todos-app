package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kocannn/todos-app/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", db_user, db_pass, db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&domain.User{}, &domain.Todo{})

	log.Println("Database is connected")

	return db

}
