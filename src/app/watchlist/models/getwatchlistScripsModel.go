package models

type Scrip struct {
	Symbol string `json:"symbol"`
	Token  uint   `json:"token"`
}
type GetWatchlistScrips struct {
	WatchlistScrip []Scrip `json:"watchlist"`
}
