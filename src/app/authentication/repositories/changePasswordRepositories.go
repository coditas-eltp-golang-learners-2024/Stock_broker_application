package repositories

import (
	"authentication/models"
	"gorm.io/gorm"
)

type AuthenticationProvider interface {
	CheckEmailAndPassword(email, password string) bool
	SetNewPassword(email, newPassword string) bool
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(dataBase *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: dataBase}
}

func (userRepository *UserDBRepository) CheckEmailAndPassword(email, password string) bool {
	var count int64
	if err := userRepository.db.Model(&models.ChangePassword{}).Where("email = ? AND password = ?", email, password).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) SetNewPassword(email, newPassword string) bool {
	var count int64
	if err := userRepository.db.Model(&models.ChangePassword{}).Where("email = ?", email).Update("password", newPassword).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
