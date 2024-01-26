package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"errors"
)

// parent struct to implement interface binding
type userService struct {
	repo domain.IUserRepo
}

// interface binding
func UserServiceInstance(userRepo domain.IUserRepo) domain.IUserService {
	return &userService{
		repo: userRepo,
	}
}

// all methods of interface are implemented
func (service *userService) GetUsers(userID uint) ([]types.UserRequest, error) {
	var allUsers []types.UserRequest
	user := service.repo.GetUsers(userID)
	if len(user) == 0 {
		return nil, errors.New("no user found")
	}
	for _, val := range user {
		allUsers = append(allUsers, types.UserRequest{
			ID:       val.ID,
			Username: val.Username,
			Password: val.Password,
		})
	}
	return allUsers, nil
}

func (service *userService) GetUsersByUsername(userName string) ([]types.UserRequest, error) {
	var allUsers []types.UserRequest
	user := service.repo.GetUsersByUsername(userName)
	if len(user) == 0 {
		return nil, errors.New("no user found")
	}
	for _, val := range user {
		allUsers = append(allUsers, types.UserRequest{
			ID:       val.ID,
			Username: val.Username,
			Password: val.Password,
		})
	}
	return allUsers, nil
}

func (service *userService) CreateUser(user *models.UserDetail) error {
	if err := service.repo.CreateUser(user); err != nil {
		return errors.New("UserDetail was not created")
	}
	return nil
}

func (service *userService) UpdateUser(user *models.UserDetail) error {
	if err := service.repo.UpdateUser(user); err != nil {
		return errors.New("UserDetail update was unsuccessful")
	}
	return nil
}
func (service *userService) DeleteUser(userID uint) error {
	if err := service.repo.DeleteUser(userID); err != nil {
		return errors.New("UserDetail deletion was unsuccessful")
	}
	return nil
}
