package repositories

import (
	"authentication/models"
	"log"

	"gorm.io/gorm"
)

type UserSignUpRepository interface {
	CheckUserExists(client *gorm.DB, user *models.User) (int64, error)
	InsertUserIntoDB(client *gorm.DB, user *models.User) error
}

type userSignUpRepository struct{}

func NewUserSignUpInstance() *userSignUpRepository {
	return &userSignUpRepository{}
}

func (repo *userSignUpRepository) CheckUserExists(db *gorm.DB, user *models.User) (int64, error) {

	var count int64
	err := db.Model(&models.User{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (repo *userSignUpRepository) InsertUserIntoDB(db *gorm.DB, user *models.User) error {

	log.Println("INFO: SignUpRepository - Inserting user into db")
	err := db.Model(&models.User{}).Create(user).Error
	if err != nil {
		log.Println("ERROR: SignUpRepository - Error inserting user into db", err)
		return err
	}
	log.Println("INFO: SignUpRepository - User inserted into db successfully")
	return nil
}
