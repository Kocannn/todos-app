package service

import "github.com/kocannn/todos-app/domain"

type userService struct {
	userRepo domain.UserRepository
}

// CreateUser implements domain.UserService.
func (u *userService) CreateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
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
