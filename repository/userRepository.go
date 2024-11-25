package repository

import (
	"github.com/kocannn/todos-app/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// CreateUser implements domain.UserRepository.
func (u *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}


func (u *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}


func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
