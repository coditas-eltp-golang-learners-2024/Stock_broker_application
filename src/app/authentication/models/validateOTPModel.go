package models

type ValidateOTPRequest struct {
	UserID uint16 `json:"id" validate:"required"`
	OTP    uint16 `json:"otp" validate:"required,numeric,min=1000,max=9999"`
}
