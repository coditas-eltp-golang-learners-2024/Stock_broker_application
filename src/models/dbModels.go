package models

import (
	"stock_broker_application/src/constants"
	"time"
)

type Users struct {
    Id             int       `gorm:"column:id" json:"id"`
    UserName       string    `gorm:"column:username" json:"username"`
    Name           string    `gorm:"column:name" json:"name"`
    Email          string    `gorm:"column:email" json:"email"`
    PhoneNumber    string    `gorm:"column:phone_number" json:"phone_number"`
    PanCard        string    `gorm:"column:pan_card" json:"pan_card"`
    Password       string    `gorm:"column:password" json:"password"`
    Token          string    `gorm:"column:token" json:"token"`
    CreationTime   time.Time `gorm:"column:created_at" json:"created_at"`
    OTP            int       `gorm:"column:otp" json:"otp"`
    EpochTimestamp int64     `gorm:"column:epochtimestamp" json:"epochtimestamp"`
}

func (Users) TableName() string {
    return constants.UserTable
}


