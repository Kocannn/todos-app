package service

import (
	"errors"

	"github.com/kocannn/todos-app/domain"
	"github.com/kocannn/todos-app/helper"
)

type userService struct {
	userRepo domain.UserRepository
}

// CreateUser implements domain.UserService.
func (u *userService) CreateUser(user *domain.User) (*domain.User, error) {

	if err := validateInput(*user); err != nil {
		return user, err
	}

	hashedPassword, err := helper.HassPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPassword

	return u.userRepo.CreateUser(user)

}

// Login implements domain.UserService.
func (u *userService) Login(email string, password string) (string, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	isPasswordValid, err := helper.ComparePassword(user.Password, password)
	if err != nil {
		return "", err
	}
	if !isPasswordValid {
		return "", errors.New("Invalid password")
	}

	token, err := helper.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil

}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func validateInput(user domain.User) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if user.Password == "" {
		return errors.New("Password is required")
	}
	return nil
}
