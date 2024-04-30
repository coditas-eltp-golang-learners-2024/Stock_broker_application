package repositories

import (
	dbModels "stock_broker_application/src/models"

	"gorm.io/gorm"
)

type ForgotPasswordRequest interface {
	VerifyCredentialsAndUpdateOTP(email string, pancardNumber string, newPassword string) error
}

type UpdatePasswordRepo struct {
	DB *gorm.DB
}

func NewforgotPasswordRepository(db *gorm.DB) ForgotPasswordRequest {
	return &UpdatePasswordRepo{DB: db}
}

func (repo *UpdatePasswordRepo) VerifyCredentialsAndUpdateOTP(email string, pancardNumber string, newPassword string) error {
	// Update the password for the user if found
	result := repo.DB.Model(&dbModels.Users{}).
		Where("email = ? AND pan_card = ?", email, pancardNumber).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		// No user found with the provided credentials
		return gorm.ErrRecordNotFound
	}

	return nil
}
