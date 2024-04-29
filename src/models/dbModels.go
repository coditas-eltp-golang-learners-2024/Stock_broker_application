package models

import (
	"github.com/dgrijalva/jwt-go"
	"stock_broker_application/src/constants"
)

type PasswordUpdateRequest struct {
	OldPassword string `gorm:"column:password" json:"oldPassword"`
	NewPassword string `gorm:"column:newPassword" json:"newPassword"`
}

func (PasswordUpdateRequest) TableName() string {
	return constants.UserTable
}

type TokenModel struct {
	Email string `gorm:"column:email" json:"email"`
	jwt.StandardClaims
}

func (TokenModel) TableName() string {
	return constants.UserTable
}
