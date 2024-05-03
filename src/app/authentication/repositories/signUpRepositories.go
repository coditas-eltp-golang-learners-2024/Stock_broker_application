package repositories

import (
	"authentication/models"
	"log"
	dbModels "stock_broker_application/src/models"

	"gorm.io/gorm"
)

type UserSignUpRepository interface {
	CheckUserExists(client *gorm.DB, user *models.UserSignUp) (int64, error)
	InsertUserIntoDB(client *gorm.DB, user *models.UserSignUp) error
}

type userSignUpRepository struct{}

func NewUserSignUpInstance() *userSignUpRepository {
	return &userSignUpRepository{}
}

func (repo *userSignUpRepository) CheckUserExists(db *gorm.DB, user *models.UserSignUp) (int64, error) {

	var count int64
	err := db.Model(&dbModels.Users{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (repo *userSignUpRepository) InsertUserIntoDB(db *gorm.DB, user *models.UserSignUp) error {

	log.Println("INFO: SignUpRepository - Inserting user into db")
	err := db.Model(&models.UserSignUp{}).Create(user).Error
	if err != nil {
		log.Println("ERROR: SignUpRepository - Error inserting user into db", err)
		return err
	}
	log.Println("INFO: SignUpRepository - User inserted into db successfully")
	return nil
}
