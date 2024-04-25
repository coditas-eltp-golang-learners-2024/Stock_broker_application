package repositories

import (
	"authentication/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// CustomerRepository defines the interface for interacting with customer data.
type CustomerRepository interface {
	AssignOtpToEmail(email string, otp int, creatime time.Time) bool
	CheckOtp(email string, otp int) bool
	UpdateUserToken(email, token string) error
}

// CustomerDBRepository is an implementation of CustomerRepository using GORM for database interactions.
type CustomerDBRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new instance of CustomerDBRepository.
func NewCustomerRepository(db *gorm.DB) *CustomerDBRepository {
	return &CustomerDBRepository{db: db}
}

func (repo *CustomerDBRepository) AssignOtpToEmail(email string, otp int, creationTime time.Time) bool {
	var count int64
	// Truncate milliseconds from otpExpiry
	creationTime = creationTime.Truncate(time.Second)
	if err := repo.db.Model(&models.OTPValidationRequest{}).Where("email = ?", email).Updates(models.OTPValidationRequest{OTP: otp, CreationTime: creationTime}).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) CheckOtp(email string, otp int) bool {
	var count int64
	var otpCreationTime mysql.NullTime

	if repo.db == nil {
		log.Printf("Database connection is nil")
		return false
	}

	err := repo.db.Table("users").Select("createdAt").Where("email = ?", email).Scan(&otpCreationTime).Error
	if err != nil {
		return false
	}
	if otpCreationTime.Valid {
		duration := time.Since(otpCreationTime.Time)
		if duration > 12000*time.Minute {
			return false
		}
	}
	if err := repo.db.Model(&models.OTPValidationRequest{}).Where("email=? AND otp =?", email, otp).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) UpdateUserToken(email, token string) error {
	// Update or insert user's token into the database
	// Example: Using GORM to update or insert token into the users table
	result := repo.db.Model(&models.OTPValidationRequest{}).Where("email = ?", email).Update("token", token)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
