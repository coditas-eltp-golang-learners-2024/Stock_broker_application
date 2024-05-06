package models

import (
	"stock_broker_application/src/constants"
	"time"
)

type Users struct {
	Id             int       `gorm:"primary_key;auto_increment" json:"id"`
	UserName       string    `gorm:"column:username" json:"username"`
	Name           string    `gorm:"column:name" json:"name"`
	Email          string    `gorm:"column:email" json:"email"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`
	PhoneNumber    uint64    `gorm:"column:phone_number" json:"phoneNumber"`
	PanCard        string    `gorm:"column:pan_card" json:"panCard"`
	Password       string    `gorm:"column:password" json:"password"`
	Token          string    `gorm:"column:token" json:"token"`
	CreationTime   time.Time `gorm:"column:created_at" json:"created_at"`
	OTP            uint16    `gorm:"column:otp" json:"otp"`
	EpochTimestamp int64     `gorm:"column:epochtimestamp" json:"epochtimestamp"`
}

type Watchlist struct {
	UserId        uint64 `json:"userId" validate:"required" gorm:"column:user_id"`
	WatchlistName string `json:"watchlistName" validate:"required" gorm:"column:watchlist_name"`
	StockId       uint64 `json:"stockId" validate:"required" gorm:"column:stock_id"`
}

func (Users) TableName() string {
	return constants.UserTable
}

func (*Watchlist) TableName() string {
	return constants.WatchlistTable
}
