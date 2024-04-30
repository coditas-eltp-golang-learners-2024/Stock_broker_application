package models

type GetWatchListRequest struct {
	UserId        int    `json:"userId"`
	WatchListName string `json:"watchListName"`
}
