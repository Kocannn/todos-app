package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kocannn/todos-app/domain"
)

func GenerateJWT(user *domain.User) (string, error) {
	type customUser struct {
		Email    string `json:"email"`
	}

	var newUsers customUser
	newUsers.Email = user.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    newUsers.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("SECRET")
	if secret == "" {
		return "", errors.New("SECRET is not set")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
