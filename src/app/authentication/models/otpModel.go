package models

import "time"

type Users struct {
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	PanCard        string    `json:"pan_card"`
	Password       string    `json:"password"`
	Token          string    `json:"token"`
	CreationTime   time.Time `json:"created_at"`
	OTP            int       `json:"otp"`
	EpochTimestamp int64     `json:"epochtimestamp"`
}
