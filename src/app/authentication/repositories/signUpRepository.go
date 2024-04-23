package repositories

import (
	"authentication/models"

	"gorm.io/gorm"
)

type UserSignUpRepository interface {
	CheckUserExists(user *models.User) (int64, error)
	InsertUserIntoDB(user *models.User) error
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserSignUpInstance(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{
		db: db,
	}
}

func (repo UserDBRepository) CheckUserExists(user *models.User) (int64, error) {

	var count int64
	err := repo.db.Model(&models.User{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (repo UserDBRepository) InsertUserIntoDB(user *models.User) error {

	err := repo.db.Model(&models.User{}).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
