package models

type ValidateOTPRequest struct {
	UserID string `json:"id" validate:"required,min=6,max=50"`
	OTP    uint16 `json:"otp" validate:"required,numeric,min=1000,max=9999"`
}
