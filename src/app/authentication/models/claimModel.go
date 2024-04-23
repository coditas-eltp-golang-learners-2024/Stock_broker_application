package models

import "github.com/golang-jwt/jwt"

type Claim struct {
	Email string `gorm:"column:email" json:"email" example:"john.doe@gmail.com"`
	Token string `gorm:"column:token" json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNhbmtldHZjaGRkYUBnbWFpbC5jb20i"`
	jwt.StandardClaims
}

func (Claim) TableName() string {
	return "users"
}
