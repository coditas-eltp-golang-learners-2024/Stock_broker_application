package models

import "time"

type OTPValidationRequest struct {
	Email        string    `gorm:"column:email" json:"email"`
	OTP          int       `gorm:"column:otp" json:"otp"`
	CreationTime time.Time `gorm:"column:createdAt" json:"createdAt"`
}

func (OTPValidationRequest) TableName() string {
	return "users"
}
