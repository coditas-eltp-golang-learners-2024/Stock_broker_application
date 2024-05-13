package handler

import (
	"net/http"
	"stock_broker_application/src/utils/validations"
	"watchlist/business"
	"watchlist/commons/constants"
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

// DeleteWatchList godoc
// @Summary Delete a watchlist
// @Description Delete a watchlist with the provided details
// @Tags DeleteWatchlist
// @Accept json
// @Produce json
// @Security JWT
// @Param watchlist body models.WatchlistDeleteModel true "Watchlist details"
// @Success 200 {string} string "Watchlist deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Router /v1/delete-watchlist [delete]
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
	ctx.JSON(http.StatusOK, gin.H{"message": constants.WatchlistDeletedSuccessMessage})
}
