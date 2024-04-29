package models

type ForgotPasswordRequest struct {
	Email         string `json:"email" example:"testUser@gmail.com"`
	PanCardNumber string `json:"pancardNumber" example:"abgjhi6789"`
	Password      string `json:"NewPassword" validate:"required,min=8,max=20,PasswordValidation" example:"sample11110"`
}

// TableName sets the table name for UserInfo explicitly.
func (ForgotPasswordRequest) TableName() string {
	return "users" // Specify the desired table name here
}
