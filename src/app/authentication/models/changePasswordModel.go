package models

type ChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required,passwordValidation" example:"S@nket123"`
	NewPassword string `json:"newPassword" validate:"required,passwordValidation" example:"Coditas@18"`
}
