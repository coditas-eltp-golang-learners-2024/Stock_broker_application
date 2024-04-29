package models

import "stock_broker_application/src/constants"

type ForgotPasswordRequest struct {
	Email         string `gorm:"column:email" json:"email"`
	PanCardNumber string `gorm:"column:pan_card" json:"pancardNumber"`
	Password      string `gorm:"column:password" json:"NewPassword"`
}

func (ForgotPasswordRequest) TableName() string {
	return constants.UserTable
}
