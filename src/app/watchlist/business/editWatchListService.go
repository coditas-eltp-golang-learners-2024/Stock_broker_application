package business

import (
	"stock_broker_application/src/utils/postgres"
	"watchlist/models"
	"watchlist/repositories"
)

type EditWatchListService struct {
	EditWatchListRepository repositories.EditWatchListRepository
}

func NewEditWatchListService(editWatchListRepository repositories.EditWatchListRepository) *EditWatchListService {
	return &EditWatchListService{
		EditWatchListRepository: editWatchListRepository,
	}
}

func (service *EditWatchListService) EditWatchList(watchlist *models.WatchlistRenameModel) error {
	client := postgres.GetPostGresClient()

	err := service.EditWatchListRepository.RenameWatchlist(client.GormDb, watchlist)
	if err != nil {
		return err
	}
	return nil

}
