package models

type GetWatchListsRequest struct {
	UserId int `json:"userID" validate:"required" example:"5" `
}
