package models

type WatchlistDeleteModel struct {
	WatchlistName string `json:"watchListName" validate:"required,min=1"`
}
