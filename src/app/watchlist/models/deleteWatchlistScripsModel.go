package models

type DeleteWatchlistScripsRequest struct {
	WatchlistName string `json:"watchlistName" validate:"required" example:"Mid watchlists"`
	Scrips        []int  `json:"scrips" validate:"required,min=1,dive,gt=0" example:"1,3"`
}
