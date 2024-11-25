package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        int       `json:"id" gorm:"primary_key;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;not null;autoUpdateTime"`
	Todos     []Todo    `json:"todos"`
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

type UserService interface {
	CreateUser(user *User) (*User, error)
	Login(email, password string) (string, error)
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
}
