package models

type ChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required" example:"S@nket123"`
	NewPassword string `json:"" validate:"required,validatePassword" example:"Coditas@18"`
}
