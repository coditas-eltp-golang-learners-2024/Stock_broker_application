package models

type RenameWatchlistRequest struct {
	WatchlistName    string `json:"watchlistName" validate:"required"`
	NewWatchlistName string `json:"newWatchlistName" validate:"required"`
}
