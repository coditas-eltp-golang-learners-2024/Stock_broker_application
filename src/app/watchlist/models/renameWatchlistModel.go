package models

type RenameWatchlist struct {
	WatchlistName    string `json:"watchlistName" validate:"required"`
	NewWatchlistName string `json:"newWatchlistName" validate:"required"`
}
