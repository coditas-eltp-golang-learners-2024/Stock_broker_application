package repositories

import (
	dbModels "stock_broker_application/src/models"

	genericConstants "stock_broker_application/src/constants"

	"gorm.io/gorm"
)

type ForgotPasswordRepository interface {
	VerifyAndUpdatePassword(email string, pancardNumber string, newPassword string) error
}

type userDBRepository struct {
	DB *gorm.DB
}

func NewForgotPasswordRepository(db *gorm.DB) ForgotPasswordRepository {
	return &userDBRepository{DB: db}
}

func (repository *userDBRepository) VerifyAndUpdatePassword(email string, pancardNumber string, newPassword string) error {
	// Update the password for the user if found
	result := repository.DB.Model(&dbModels.Users{}).
		Where("email = ? AND pan_card = ?", email, pancardNumber).
		Update(genericConstants.Password, newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
