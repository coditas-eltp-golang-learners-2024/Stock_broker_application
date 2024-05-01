package repositories

import (
	"errors"
	"gorm.io/gorm"
	dbModels "stock_broker_application/src/models"
)

type SignInRepository interface {
	AuthenticateUser(username string, password string) (bool, error)
}

type SignInRepositoryImpl struct {
	db *gorm.DB
}

// NewSignInRepositoryImpl creates a new instance of SignInRepositoryImpl
func NewSignInRepositoryImpl(db *gorm.DB) *SignInRepositoryImpl {
	return &SignInRepositoryImpl{db: db}
}

func (repo *SignInRepositoryImpl) AuthenticateUser(username string, password string) (bool, error) {
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
