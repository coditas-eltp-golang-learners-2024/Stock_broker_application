package models

type Watchlist struct {
	UserId        uint64 `json:"userId" validate:"required" gorm:"column:user_id"`
	WatchlistName string `json:"watchlistName" validate:"required" gorm:"column:watchlist_name"`
	StockId       uint64 `json:"stockId" validate:"required" gorm:"column:stock_id"`
}

func (*Watchlist) TableName() string {
	return "watchlist"
}
