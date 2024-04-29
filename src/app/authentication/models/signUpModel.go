package models

type User struct {
	Name        string `json:"name" validate:"required" gorm:"column:name"`
	Email       string `json:"email" validate:"required,email" gorm:"column:email"`
	PhoneNumber uint64 `json:"phoneNumber" validate:"required,min=1000000000,max=9999999999,numeric" gorm:"column:phone_number"`
	PanCard     string `json:"panCard" validate:"required" gorm:"column:pan_card"`
	Password    string `json:"password" validate:"required,min=8,max=20" gorm:"column:password"`
}

func (*User) TableName() string {
	return "users"
}
