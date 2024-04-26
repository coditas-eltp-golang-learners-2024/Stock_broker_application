package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	Email string `gorm:"column:email" json:"email" example:"john.doe@gmail.com"`
	jwt.StandardClaims
}

func (Claim) TableName() string {
	return "users"
}