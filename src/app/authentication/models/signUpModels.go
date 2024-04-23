package models

type User struct {
	Name        string `json:"name" validate:"required" gorm:"column:name"`
	Email       string `json:"email" validate:"required,email" gorm:"column:email"`
	PhoneNumber int    `json:"phoneNumber" validate:"required" gorm:"column:phoneNumber"`
	PanCard     string `json:"panCard" validate:"required" gorm:"column:panCard"`
	Password    string `json:"password" validate:"required,passwordValidation" gorm:"column:password"`
}

func (*User) TableName() string {
	return "users"
}
