package models

type ForgotPasswordRequest struct {
	Email         string `json:"email" validate:"required,email" example:"testUser@gmail.com" `
	PanCardNumber string `json:"pancardNumber" validate:"required,len=10" example:"abgjhi6789" `
	Password      string `json:"NewPassword" validate:"required=password,min=8,max=20,PasswordValidation" example:"sample11110"`
}
