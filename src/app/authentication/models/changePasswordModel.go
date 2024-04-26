package models

type ChangePassword struct {
	OldPassword string `gorm:"column:password" json:"oldPassword" validate:"required,min=8" example:"password"`
	NewPassword string `gorm:"column:newPassword" json:"newPassword" validate:"required,min=8" example:"newPassword"`
}

func (ChangePassword) TableName() string {
	return "users"
}
