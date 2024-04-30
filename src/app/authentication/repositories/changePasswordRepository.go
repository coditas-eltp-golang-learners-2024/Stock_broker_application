package repositories

import (
	"authentication/commons/constants"
	"gorm.io/gorm"
	"stock_broker_application/src/models"
)

type AuthenticationProvider interface {
	CheckEmailAndPassword(condition map[string]interface{}) bool
	SetNewPassword(PasswordUpdateSQLCondition map[string]interface{}) bool
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(dataBase *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: dataBase}
}

func (userRepository *UserDBRepository) CheckEmailAndPassword(condition map[string]interface{}) bool {
	var count int64
	if err := userRepository.db.Model(&models.Users{}).Where(condition).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) SetNewPassword(PasswordUpdateSQLCondition map[string]interface{}) bool {
	var count int64
	if err := userRepository.db.Model(&models.Users{}).Where(constants.UserName, PasswordUpdateSQLCondition[constants.UserName]).Update(constants.Password, PasswordUpdateSQLCondition[constants.Password]).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
