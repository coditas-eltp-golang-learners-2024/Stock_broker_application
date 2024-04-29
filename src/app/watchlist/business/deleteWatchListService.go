package business

import (
	"stock_broker_application/src/utils/postgres"
	"watchlist/models"
	"watchlist/repositories"
)

type DeleteWatchListService struct {
	deleteWatchListRepository repositories.DeleteWatchListRepository
}

func NewDeleteWatchListService(deleteWatchListRepository repositories.DeleteWatchListRepository) *DeleteWatchListService {
	return &DeleteWatchListService{
		deleteWatchListRepository: deleteWatchListRepository,
	}
}

func (service *DeleteWatchListService) DeleteWatchList(watchlist *models.WatchlistDeleteModel) error {

	client := postgres.GetPostGresClient()
	err := service.deleteWatchListRepository.DeleteWatchlist(client.GormDb, watchlist)
	if err != nil {
		return err
	}
	return nil
}
