package models

type CreateWatchlist struct {
	WatchlistName string `json:"watchlistName" validate:"required" example:"Mid Watchlist"`
	Scrips        []int  `json:"scrips" validate:"required" example:"44,22"`
}
