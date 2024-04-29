package handler

import (
	"net/http"
	"watchlist/business"
	"watchlist/models"

	"github.com/gin-gonic/gin"
)

type editWatchListController struct {
	service *business.EditWatchListService
}

func NewEditWatchListController(service *business.EditWatchListService) *editWatchListController {
	return &editWatchListController{
		service: service,
	}
}
func (controller *editWatchListController) EditWatchList(ctx *gin.Context) {
	var watchlist models.WatchlistRenameModel
	if err := ctx.ShouldBindJSON(&watchlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.service.EditWatchList(&watchlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Watchlist renamed successfully"})
}
