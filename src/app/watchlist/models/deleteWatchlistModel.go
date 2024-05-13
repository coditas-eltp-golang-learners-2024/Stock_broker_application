package models

type DeleteWatchlistScripsRequest struct {
	WatchlistName string `json:"watchlist_name" validate:"required" example:"Mid watchists"`
	Scrips        []int  `json:"scrips" validate:"required" example:"1,3"`
}
