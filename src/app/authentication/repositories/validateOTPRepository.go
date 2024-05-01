package repositories

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"stock_broker_application/src/constants"
	"stock_broker_application/src/models"
	"time"
)

type CustomerRepository interface {
	CheckOtp(username string, otp uint16) bool
	UpdateUserToken(username, token string) error
}

type CustomerDBRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new instance of CustomerDBRepository.
func NewCustomerRepository(db *gorm.DB) *CustomerDBRepository {
	return &CustomerDBRepository{db: db}
}

func (repo *CustomerDBRepository) CheckOtp(username string, otp uint16) bool {
	var count int64
	var otpCreationTime mysql.NullTime

	if repo.db == nil {
		log.Printf("Database connection is nil")
		return false
	}

	err := repo.db.Table(constants.UserTable).Select("createdAt").Where("username = ?", username).Scan(&otpCreationTime).Error
	if err != nil {
		return false
	}
	if otpCreationTime.Valid {
		duration := time.Since(otpCreationTime.Time)
		if duration > 12000*time.Minute {
			return false
		}
	}
	if err := repo.db.Model(&models.Users{}).Where("username=? AND otp =?", username, otp).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) UpdateUserToken(username, token string) error {
	result := repo.db.Model(&models.Users{}).Where("username = ?", username).Update("token", token)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
