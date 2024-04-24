package models


type SignInRequest struct {
	Email    string `gorm:"column:email" json:"email" validate:"required,email" example:"john.doe@gmail.com"`
	Password string `gorm:"column:password" json:"password" validate:"required,min=8" example:"password"`
}

//  name of the database table for the SignInRequest model
func (SignInRequest) TableName() string {
	return "users"
}
