package repositories

import (
	"gorm.io/gorm"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"time"
)

type UserRepository interface {
	UpdateOTPAndCreationTime(email string, newOTP int) error
}

type UserDBRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserDBRepository.
func NewUserRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (repo *UserDBRepository) UpdateOTPAndCreationTime(email string, newOTP int) error {
	if err := repo.db.Model(&models.Users{}).Where("username = ?", email).Update(constants.OTP, newOTP).Error; err != nil {
		return err
	}
	otpCreationTime := time.Now().Truncate(time.Second)
	if err := repo.db.Model(&models.Users{}).Where("username = ?", email).Update(constants.CreatedAt, otpCreationTime).Error; err != nil {
		return err
	}
	return nil
}
