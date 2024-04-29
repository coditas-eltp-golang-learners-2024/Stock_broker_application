package models

import (
	"stock_broker_application/src/constants"
)

type PasswordUpdateRequest struct {
	OldPassword string `gorm:"column:password" json:"oldPassword"`
	NewPassword string `gorm:"column:newPassword" json:"newPassword"`
}

func (PasswordUpdateRequest) TableName() string {
	return constants.UserTable
}
