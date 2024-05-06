package repositories

import (
	"gorm.io/gorm"
	genericConstants "stock_broker_application/src/constants"
	dbModels "stock_broker_application/src/models"
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
