package models

type UserSignUp struct {
	Name        string `json:"name" validate:"required"`
	UserName    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber uint64 `json:"phoneNumber" validate:"required,min=1000000000,max=9999999999,numeric"`
	PanCard     string `json:"panCard" validate:"required"`
	Password    string `json:"password" validate:"required,PasswordValidation,min=8,max=20"`
}

func (*UserSignUp) TableName() string {
	return "users"
}
