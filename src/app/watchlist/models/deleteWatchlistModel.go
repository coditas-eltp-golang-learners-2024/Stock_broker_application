package models

type DeleteWatchlistRequest struct {
	WatchlistName string `json:"watchListName" validate:"required"`
}
