package handler

import (
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils"
	constants "watchlist/commons/constants"

	"watchlist/business"
	"watchlist/models"

	"github.com/gin-gonic/gin"
)

type editWatchListController struct {
	service *business.EditWatchlistService
}

func NewEditWatchListController(service *business.EditWatchlistService) *editWatchListController {
	return &editWatchListController{
		service: service,
	}
}

// @Summary Edit a watchlist
// @Description Edit a watchlist with the provided details
// @Tags EditWatchlist
// @Accept json
// @Produce json
// @Security JWT
// @Param watchlist body models.EditWatchlist true "Watchlist details"
// @Success 200 {string} string "Watchlist edited successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/edit-watchlist [put]
func (controller *editWatchListController) EditWatchList(ctx *gin.Context) {

	var editWatchListRequest models.EditWatchlist
	if err := ctx.ShouldBindJSON(&editWatchListRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
		return
	}
	if err := controller.service.EditWatchlist(&editWatchListRequest, ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONErrorMessage: err.Error()})
		return
	}
	utils.SendStatusOkSuccess(ctx, constants.WatchlistEditSuccessMessage)
}
