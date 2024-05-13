package handler

import (
	"net/http"
	"stock_broker_application/src/utils/validations"
	"watchlist/business"
	constants "watchlist/commons/constants"
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

// EditWatchList godoc
// @Summary Edit a watchlist
// @Description Edit a watchlist with the provided details
// @Tags EditWatchlist
// @Accept json
// @Produce json
// @Security JWT
// @Param watchlist body models.WatchlistRenameModel true "Watchlist details"
// @Success 200 {string} string "Watchlist edited successfully"
// @Failure 400 {string} string "Bad request"
// @Router /v1/edit-watchlist [put]
func (controller *editWatchListController) EditWatchList(ctx *gin.Context) {
	var watchlist models.WatchlistRenameModel
	if err := ctx.ShouldBindJSON(&watchlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validations.GetCustomValidator(ctx.Request.Context()).Struct(watchlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := controller.service.EditWatchList(&watchlist, ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": constants.WatchlistRenameSuccessMessage})
}
