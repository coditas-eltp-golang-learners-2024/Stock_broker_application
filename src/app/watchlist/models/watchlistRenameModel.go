package models

type WatchlistRenameModel struct {
	WatchlistName    string `json:"watchlistName" validate:"required"`
	NewWatchlistName string `json:"newWatchlistName" validate:"required,min=1"`
}
