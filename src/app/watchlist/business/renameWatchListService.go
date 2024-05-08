package business

import (
	genericConstants "stock_broker_application/src/constants"
	"stock_broker_application/src/utils/postgres"
	"watchlist/models"
	"watchlist/repositories"

	"github.com/gin-gonic/gin"
)

type EditWatchListService struct {
	EditWatchListRepository repositories.EditWatchListRepository
}

func NewEditWatchListService(editWatchListRepository repositories.EditWatchListRepository) *EditWatchListService {
	return &EditWatchListService{
		EditWatchListRepository: editWatchListRepository,
	}
}

func (service *EditWatchListService) EditWatchList(watchlist *models.WatchlistRenameModel, ctx *gin.Context) error {
	client := postgres.GetPostGresClient()
	userId := ctx.GetString(genericConstants.Id)

	err := service.EditWatchListRepository.RenameWatchlist(client.GormDb, watchlist, userId)
	if err != nil {
		return err
	}
	return nil

}
