package service

import (
	"errors"

	"github.com/kocannn/todos-app/domain"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo domain.UserRepository
}

// CreateUser implements domain.UserService.
func (u *userService) CreateUser(user domain.User) (domain.User, error) {

	if err := validateInput(user); err != nil {
		return user, err
	}
	pw := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	if err != nil {
		return user, err
	}

	return u.userRepo.CreateUser(user)

}

// Login implements domain.UserService.
func (u *userService) Login(email string, password string) (domain.User, error) {
	panic("unimplemented")
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
