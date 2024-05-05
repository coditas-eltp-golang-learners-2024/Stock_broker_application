package models

type ChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required" example:"S@nket123"`
	NewPassword string `json:"newPassword" validate:"required,min=8,max=20,PasswordValidation" example:"Coditas@18"`
}
