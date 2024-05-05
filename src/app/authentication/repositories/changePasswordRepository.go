package repositories

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"gorm.io/gorm"
)

type ChangePasswordRepositor interface {
	CheckUsernameAndPassword(condition map[string]interface{}) bool
	SetNewPassword(PasswordUpdateSQLCondition map[string]interface{}) bool
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(dataBase *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: dataBase}
}

func (userRepository *UserDBRepository) CheckUsernameAndPassword(condition map[string]interface{}) bool {
	var count int64
	if err := userRepository.db.Model(&models.Users{}).Where(condition).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (userRepository *UserDBRepository) SetNewPassword(PasswordUpdateSQLCondition map[string]interface{}) bool {
	var count int64
	if err := userRepository.db.Model(&models.Users{}).Where(genericConstants.Username, PasswordUpdateSQLCondition[genericConstants.Username]).Updates(PasswordUpdateSQLCondition).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
