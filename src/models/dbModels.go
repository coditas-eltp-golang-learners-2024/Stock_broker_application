package models

import "stock_broker_application/src/constants"

type Watchlist struct {
	ID            int    `gorm:"column:id" json:"id"`
	UserID        int    `gorm:"column:user_id" json:"userID"`
	WatchlistName string `gorm:"column:watchlist_name" json:"watchlistName"`
}

func (Watchlist) TableName() string {
	return constants.Watchlist
}
