package repositories

import (
	"errors"
	"gorm.io/gorm"
	genericConstants"stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	dbModels "stock_broker_application/src/models"
	"time"
)

type SignInRepository interface {
	AuthenticateUser(username string, password string) (bool, error)
	UpdateOTPAndCreationTime(email string, newOTP int) error
}

type userSignInDBRepository struct {
	db *gorm.DB
}

// NewSignInRepositoryImpl creates a new instance of SignInRepositoryImpl
func NewSignInRepository(db *gorm.DB) *userSignInDBRepository {
	return &userSignInDBRepository{db: db}
}

func (repo *userSignInDBRepository) AuthenticateUser(username string, password string) (bool, error) {
	var user dbModels.Users
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

func (repo *userSignInDBRepository) UpdateOTPAndCreationTime(email string, newOTP int) error {
	if err := repo.db.Model(&models.Users{}).Where("username = ?", email).Update(genericConstants.OTP, newOTP).Error; err != nil {
		return err
	}
	otpCreationTime := time.Now().Truncate(time.Second)
	if err := repo.db.Model(&models.Users{}).Where("username = ?", email).Update(genericConstants.CreatedAt, otpCreationTime).Error; err != nil {
		return err
	}
	epochTime := time.Now().Unix()
	if err := repo.db.Model(&models.Users{}).Where("username = ?", email).Update(genericConstants.EpochTimestamp, epochTime).Error; err != nil {
		return err
	}
	return nil
}