package models

type SignInRequest struct {
	UserName string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=20,password"`
}
