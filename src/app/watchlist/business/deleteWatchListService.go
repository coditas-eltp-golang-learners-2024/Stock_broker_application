package business

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"
	"watchlist/models"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type DeleteWatchListService struct {
	deleteWatchListRepository repositories.DeleteWatchListRepository
}

func NewDeleteWatchListService(deleteWatchListRepository repositories.DeleteWatchListRepository) *DeleteWatchListService {
	return &DeleteWatchListService{
		deleteWatchListRepository: deleteWatchListRepository,
	}
}

func (service *DeleteWatchListService) DeleteWatchList(watchlist *models.DeleteWatchlistRequest, ctx *gin.Context) error {

	client := postgres.GetPostGresClient()
	userId := ctx.Value(genericConstants.Id).(uint16)
	err := service.deleteWatchListRepository.DeleteWatchlist(client.GormDb, watchlist, userId)
	if err != nil {
		return err
	}
	return nil
}
