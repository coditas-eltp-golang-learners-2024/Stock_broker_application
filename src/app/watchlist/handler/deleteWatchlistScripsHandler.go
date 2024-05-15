package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
	"watchlist/commons/constants"
	"watchlist/models"
)

type DeleteWatchlistScripsController struct {
	service *business.DeleteWatchlistScripsService
}

func NewDeleteWatchlistScripsController(service *business.DeleteWatchlistScripsService) *DeleteWatchlistScripsController {
	return &DeleteWatchlistScripsController{
		service: service,
	}
}

// HandleDeleteWatchlistScrips handles HTTP DELETE requests for deleting scrips from a watchlist.
// @Summary Delete scrips from a watchlist
// @Description Delete scrips from a watchlist based on the provided watchlist name and scrips
// @Accept  json
// @Produce  json
// @Param request body models.DeleteWatchlistScripsRequest true "Delete Watchlist Scrips Request"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to delete scrips"
// @Router /v1/watchlist/scrips [delete]
func (controller *DeleteWatchlistScripsController) HandleDeleteWatchlistScrips(context *gin.Context) {
	var request models.DeleteWatchlistScripsRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{genericConstants.GenericJSONMessage: constants.InvalidRequestPayloadError})
		return
	}

	err := controller.service.DeleteScripsFromWatchlist(context, request.WatchlistName, request.Scrips)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{genericConstants.GenericJSONMessage: constants.FailedToDeleteScripsError})
		return
	}

	context.JSON(http.StatusOK, gin.H{genericConstants.GenericJSONMessage: constants.DeletedWatchlistScripsSuccessMessage})
}
