package models

type User struct {
	Name        string `json:"name" validate:"required" gorm:"column:name"`
	Email       string `json:"email" validate:"required,email" gorm:"column:email"`
	PhoneNumber int    `json:"phoneNumber" validate:"required,min=1000000000,max=9999999999,numeric" gorm:"column:phoneNumber"`
	PanCard     string `json:"panCard" validate:"required" gorm:"column:panCard"`
	Password    string `json:"password" validate:"required,min=8,max=20,passwordValidation" gorm:"column:password"`
}

func (*User) TableName() string {
	return "users"
}
