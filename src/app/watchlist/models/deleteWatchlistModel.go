package models

type DeleteWatchlistScripsRequest struct {
	WatchlistName string `json:"watchlist_name" validate:"required" example:"Mid Watchlist"`
	Scrips        []int  `json:"scrips" validate:"required" example:"44,22"`
}
