package models

type SignInRequest struct {
	UserName string `json:"username" validate:"required" example:"virat"`
	Password string `json:"password" validate:"required,min=8,max=20,custompassword" example:"Goat@018"`
}
