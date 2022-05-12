package database

import (
	"github.com/seanvm/serverless-golang-starter-backend/app"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// CreateUser creates a user
func (s *UserRepository) CreateUser(user *app.User) (*app.User, error) {
	result := s.db.Create(user)

	return user, result.Error
}

// GetUser returns a user with a specific ID
func (s *UserRepository) GetUser(id string) (*app.User, error) {
	user := app.User{}
	result := s.db.Find(&user, id)

	return &user, result.Error
}
