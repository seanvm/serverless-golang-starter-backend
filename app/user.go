package app

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      *string   `json:"name"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRepository interface {
	GetUser(id string) (*User, error)
	CreateUser(user *User) (*User, error)
}

type UserServicer interface {
	GetUserRepo() UserRepository
	GetUser(id string) (*User, error)
	CreateUser(user *User) (*User, error)
}

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(ur UserRepository) UserServicer {
	return &UserService{
		UserRepository: ur,
	}
}

func (us *UserService) GetUserRepo() UserRepository {
	return us.UserRepository
}

func (us *UserService) CreateUser(user *User) (*User, error) {
	return us.GetUserRepo().CreateUser(user)
}

func (us *UserService) GetUser(id string) (*User, error) {
	return us.GetUserRepo().GetUser(id)
}
