package models

type ValidateOTPRequest struct {
	UserName string `json:"username" validate:"required"`
	OTP      uint16 `json:"otp" validate:"required"`
}
