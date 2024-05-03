package repositories

import (
	"errors"
	"gorm.io/gorm"
	"stock_broker_application/src/models"
)

type UserRepository interface {
	CheckOtp(username string, otp uint16) (bool, error)
	CheckUserExists(username string) (bool, error)
	UpdateUserToken(username, token string) error
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (repo *UserDBRepository) CheckOtp(username string, otp uint16) (bool, error) {
	var count int64
	if err := repo.db.Model(&models.Users{}).Where("username = ? AND otp = ?", username, otp).Count(&count).Error; err != nil {
		return false, errors.New("error checking OTP: " + err.Error())
	}
	return count > 0, nil
}

func (repo *UserDBRepository) CheckUserExists(username string) (bool, error) {
	var count int64
	err := repo.db.Model(&models.Users{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserDBRepository) UpdateUserToken(username, token string) error {
	result := repo.db.Model(&models.Users{}).Where("username = ?", username).Update("token", token)
	if result.Error != nil {
		return errors.New("error updating user token: " + result.Error.Error())
	}
	return nil
}
