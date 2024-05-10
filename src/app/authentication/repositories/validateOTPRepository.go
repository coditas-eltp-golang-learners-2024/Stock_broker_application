package repositories

import (
	"authentication/commons/constants"
	"errors"
	"stock_broker_application/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CheckOtp(userID uint16, otp uint16) (bool, error)
	CheckUserExists(userID uint16) (bool, error)
	UpdateUserToken(userID uint16, token string) error
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (repo *UserDBRepository) CheckOtp(userID uint16, otp uint16) (bool, error) {
	var count int64
	if err := repo.db.Model(&models.Users{}).Where("id = ? AND otp = ?", userID, otp).Count(&count).Error; err != nil {
		return false, errors.New(constants.ErrorUpdateUserToken + err.Error())
	}
	return count > 0, nil
}

func (repo *UserDBRepository) CheckUserExists(userID uint16) (bool, error) {
	var count int64
	err := repo.db.Model(&models.Users{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserDBRepository) UpdateUserToken(userID uint16, token string) error {
	result := repo.db.Model(&models.Users{}).Where("id = ?", userID).Update("token", token)
	if result.Error != nil {
		return errors.New(constants.ErrorUpdateUserToken + result.Error.Error())
	}
	return nil
}
