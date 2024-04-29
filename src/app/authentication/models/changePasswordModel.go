package models

type ChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required,min=8" example:"password"`
	NewPassword string `json:"newPassword" validate:"required,min=8" example:"newPassword"`
}
