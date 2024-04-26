package repositories

import (
	"authentication/models"

	"gorm.io/gorm"
)

type ForgetPasswordRequest interface {
	VerifyCredentialsAndUpdateOTP(email string, pancardNumber string, newPassword string) error
}

type UpdatePasswordRepo struct {
	DB *gorm.DB
}

func NewForgetPasswordRepository(db *gorm.DB) ForgetPasswordRequest {
	return &UpdatePasswordRepo{DB: db}
}

func (repo *UpdatePasswordRepo) VerifyCredentialsAndUpdateOTP(email string, pancardNumber string, newPassword string) error {
	// Update the password for the user if found
	result := repo.DB.Model(&models.ForgetPasswordRequest{}).
		Where("email = ? AND pancard = ?", email, pancardNumber).
		Update("password", newPassword)

	if result.Error != nil {
		// Handle database error
		return result.Error
	}

	// Check if any rows were affected by the update
	if result.RowsAffected == 0 {
		// No user found with the provided credentials
		return gorm.ErrRecordNotFound
	}

	return nil // Password updated successfully
}
