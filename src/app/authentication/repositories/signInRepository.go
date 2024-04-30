package repositories

import (
	"authentication/models"
	"errors"

	"gorm.io/gorm"
)

// SignInRepository defines methods for interacting with user data in the database
type SignInRepository interface {
	GetUserByUsername(username string) *models.SignInRequest
	AuthenticateUser(user *models.SignInRequest, password string) (bool, error)
}

// SignInRepositoryImpl is the implementation of SignInRepository
type SignInRepositoryImpl struct {
	db *gorm.DB
}

// NewSignInRepositoryImpl creates a new instance of SignInRepositoryImpl
func NewSignInRepositoryImpl(db *gorm.DB) *SignInRepositoryImpl {
	return &SignInRepositoryImpl{db: db}
}

func (repo *SignInRepositoryImpl) GetUserByUsername(username string) *models.SignInRequest {
	var user models.SignInRequest
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}
	return &user
}

func (repo *SignInRepositoryImpl) AuthenticateUser(user *models.SignInRequest, password string) (bool, error) {
	if user == nil {
		return false, errors.New("user not found")
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}
