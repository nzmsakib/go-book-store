package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

// for database repository operation (call from service)
type IUserRepo interface {
	GetUsers(userID uint) []models.UserDetail
	GetUsersByUsername(userName string) []models.UserDetail
	CreateUser(user *models.UserDetail) error
	UpdateUser(user *models.UserDetail) error
	DeleteUser(userID uint) error
}

// for service operation (response to controller | call from controller)
type IUserService interface {
	GetUsers(userID uint) ([]types.UserRequest, error)
	GetUsersByUsername(userName string) ([]types.UserRequest, error)
	CreateUser(user *models.UserDetail) error
	UpdateUser(user *models.UserDetail) error
	DeleteUser(userID uint) error
}
