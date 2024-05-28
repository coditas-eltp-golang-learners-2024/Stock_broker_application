package business

import (
	"errors"
	"github.com/gin-gonic/gin"
	genericConstants "stock_broker_application/src/constants"
	"watchlist/commons/constants"
	"watchlist/models"
	"watchlist/repositories"
)

type DeleteWatchlistScripsService struct {
	deleteWatchlistRepository repositories.DeleteWatchlistRepository
}

func NewDeleteWatchlistService(deleteWatchlistRepository repositories.DeleteWatchlistRepository) *DeleteWatchlistScripsService {
	return &DeleteWatchlistScripsService{
		deleteWatchlistRepository: deleteWatchlistRepository,
	}
}

func (service *DeleteWatchlistScripsService) DeleteScripsFromWatchlist(ctx *gin.Context, request models.DeleteWatchlistScripsRequest) error {
	userID := ctx.Value(genericConstants.Id).(uint16)
	watchlistCondition := map[string]interface{}{
		genericConstants.UserId:        userID,
		genericConstants.WatchlistName: request.WatchlistName,
	}

	err := service.deleteWatchlistRepository.DeleteScrips(ctx, request.WatchlistName, request.Scrips, watchlistCondition)
	if err != nil {
		return errors.New(constants.FailedToDeleteScripsfromWatchlistError)
	}

	return nil
}
