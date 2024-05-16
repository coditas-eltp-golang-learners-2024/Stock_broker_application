package models

import (
	"stock_broker_application/src/constants"
	"time"
)

type Watchlist struct {
	ID            int    `gorm:"column:id" json:"id"`
	UserID        int    `gorm:"column:user_id" json:"userID"`
	WatchlistName string `gorm:"column:watchlist_name" json:"watchlistName"`
}

func (Watchlist) TableName() string {
	return constants.WatchlistTable
}

type Users struct {
	Id             uint16    `gorm:"primary_key;auto_increment" json:"id"`
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

func (Users) TableName() string {
	return constants.UserTable
}

type Stocks struct {
	ID     uint   `gorm:"column:id;primary_key" json:"id"`
	Token  uint   `gorm:"column:token" json:"token"`
	Symbol string `gorm:"column:symbol" json:"symbol"`
}

func (Stocks) TableName() string {
	return constants.StockTable
}

type WatchlistStock struct {
	ID          uint `gorm:"column:id;primary_key" json:"id"`
	WatchlistID uint `gorm:"column:watchlist_id" json:"watchlistId"`
	StockID     uint `gorm:"column:stocks_id" json:"stockId"`
}

func (WatchlistStock) TableName() string {
	return constants.WatchlistStockTable
}
