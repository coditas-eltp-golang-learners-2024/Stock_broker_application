package models

type ForgetPasswordRequest struct {
	Email         string `json:"email" example:"testUser@gmail.com" gorm:"column:email;unique"`
	PanCardNumber string `json:"pancardNumber" example:"abgjhi6789" gorm:"column:pancard"`
	Password      string `json:"NewPassword" validate:"required,min=8,max=20,PasswordValidation" example:"sample11110" gorm:"column:password"`
}

// TableName sets the table name for UserInfo explicitly.
func (ForgetPasswordRequest) TableName() string {
	return "users" // Specify the desired table name here
}
