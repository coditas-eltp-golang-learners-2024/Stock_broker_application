package models

type DeleteWatchlist struct {
	WatchlistName string `json:"watchListName" validate:"required"`
}
