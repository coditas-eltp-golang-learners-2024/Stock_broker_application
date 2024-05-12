package models

type DeleteWatchlistScripsRequest struct {
	WatchlistName string   `json:"watchlist_name"`
	Scrips        []string `json:"scrips"`
}
