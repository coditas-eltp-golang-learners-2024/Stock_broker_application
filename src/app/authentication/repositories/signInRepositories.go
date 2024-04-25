package repositories

import (
	"authentication/models"
	"errors"

	"gorm.io/gorm"
)

// SignInRepository defines methods for interacting with user data in the database
type SignInRepository interface {
	GetUserByEmail(email string) *models.SignInRequest
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

func (r *SignInRepositoryImpl) GetUserByEmail(email string) *models.SignInRequest {
	var user models.SignInRequest
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}
	return &user
}

func (r *SignInRepositoryImpl) AuthenticateUser(user *models.SignInRequest, password string) (bool, error) {
	if user == nil {
		return false, errors.New("user not found")
	}

	// Compare the provided password with the password stored in the database
	if user.Password != password {
		// Passwords do not match
		return false, nil
	}

	// Passwords match
	return true, nil
}
