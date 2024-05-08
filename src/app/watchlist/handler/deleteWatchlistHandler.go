package handler

import (
	"net/http"
	"stock_broker_application/src/utils/validations"
	"watchlist/business"
	"watchlist/models"

	"github.com/gin-gonic/gin"
)

type deleteWatchListController struct {
	service *business.DeleteWatchListService
}

func NewEditWatchListDeleteController(service *business.DeleteWatchListService) *deleteWatchListController {
	return &deleteWatchListController{
		service: service,
	}

}

func (controller *deleteWatchListController) DeleteWatchList(ctx *gin.Context) {
	var watchlist models.WatchlistDeleteModel
	if err := ctx.ShouldBindJSON(&watchlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validations.GetCustomValidator(ctx.Request.Context()).Struct(watchlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := controller.service.DeleteWatchList(&watchlist, ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Watchlist deleted successfully"})
}
