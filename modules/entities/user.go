package entities

import (
	"time"
)

type UserUsecase interface {
	GetUsers() ([]UserResponse, error)
	GetUserByID(id int) (*UserResponse, error)
}

type UserRepository interface {
	GetUsers() ([]UserResponse, error)
	GetUserByID(id int) (*UserResponse, error)
}

type UserService interface {
	GetUsers() ([]UsersMockResponse, error)
	GetUserByID(id int) (*UsersMockResponse, error)
}

type UserResponse struct {
	UserId    int       `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	UserSince time.Time `json:"user_since" db:"user_since"`
	LastVisit time.Time `json:"last_visit" db:"last_visit"`
}
type UserResponseService struct {
	UserId   int    `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type UsersMockResponse struct {
	UserId   int    `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}
