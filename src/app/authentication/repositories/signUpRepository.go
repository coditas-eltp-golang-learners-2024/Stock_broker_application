package repositories

import (
	dbModels "stock_broker_application/src/models"
	"gorm.io/gorm"
)

type UserSignUpRepository interface {
	CheckUserExists(client *gorm.DB, user *dbModels.Users) (int64, error)
	InsertUserIntoDB(client *gorm.DB, user *dbModels.Users) error
}

type userSignUpRepository struct{}

func NewUserSignUpInstance() *userSignUpRepository {
	return &userSignUpRepository{}
}

func (repo *userSignUpRepository) CheckUserExists(db *gorm.DB, user *dbModels.Users) (int64, error) {

	var count int64
	err := db.Model(&dbModels.Users{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (repo *userSignUpRepository) InsertUserIntoDB(db *gorm.DB, user *dbModels.Users) error {

	err := db.Model(&dbModels.Users{}).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
