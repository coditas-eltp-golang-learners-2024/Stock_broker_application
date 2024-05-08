package handler

import (
	"fmt"
	"net/http"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/business"
	"watchlist/models"

	"github.com/gin-gonic/gin"
)

type WatchlistController struct {
	service *business.WatchlistService
}

func NewWatchlistController(service *business.WatchlistService) *WatchlistController {
	return &WatchlistController{
		service: service,
	}
}
func (controller *WatchlistController) HandleWatchlist(context *gin.Context) {
	var request models.WatchlistWithScripsRequest


	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := context.Value(genericConstants.Id).(uint16)
	fmt.Println(userID)
	if err := controller.service.GetScrips(context, request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, "Scrips fetched successfully")

}


