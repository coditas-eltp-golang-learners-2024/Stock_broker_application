package models

type ChangePassword struct {
	Email       string `gorm:"column:email" json:"email" validate:"required,email" example:"john.doe@gmail.com"`
	OldPassword string `gorm:"column:password" json:"oldPassword" validate:"required,min=8" example:"password"`
	NewPassword string `gorm:"column:newPassword" json:"newPassword" validate:"required,min=8" example:"newPassword"`
}

func (ChangePassword) TableName() string {
	return "users"
}
