package models

type SignInRequest struct {
	UserName string `json:"userName" validate:"required=username" example:"virat"`
	Password string `json:"password" validate:"required=password,min=8,max=20,PasswordValidation" example:"Goat@018"`
}
