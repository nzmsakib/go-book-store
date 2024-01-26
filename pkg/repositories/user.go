package repositories

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"

	"gorm.io/gorm"
)

// parent struct to implement interface binding
type userRepo struct {
	db *gorm.DB
}

// interface binding
func UserDBInstance(d *gorm.DB) domain.IUserRepo {
	return &userRepo{
		db: d,
	}
}

// all methods of interface are implemented
func (repo *userRepo) GetUsers(userID uint) []models.UserDetail {
	var user []models.UserDetail
	var err error

	if userID != 0 {
		err = repo.db.Where("id = ?", userID).Find(&user).Error
	} else {
		err = repo.db.Find(&user).Error
	}
	if err != nil {
		return []models.UserDetail{}
	}
	return user
}

func (repo *userRepo) GetUsersByUsername(userName string) []models.UserDetail {
	var user []models.UserDetail
	var err error

	if userName != "" {
		err = repo.db.Where("user_name = ?", userName).Find(&user).Error
	} else {
		err = repo.db.Find(&user).Error
	}
	if err != nil {
		return []models.UserDetail{}
	}
	return user
}

func (repo *userRepo) CreateUser(user *models.UserDetail) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *userRepo) UpdateUser(user *models.UserDetail) error {
	if err := repo.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
func (repo *userRepo) DeleteUser(userID uint) error {
	var User models.UserDetail
	if err := repo.db.Where("id = ?", userID).Delete(&User).Error; err != nil {
		return err
	}
	return nil
}
