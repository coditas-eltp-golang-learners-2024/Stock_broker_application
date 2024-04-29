package repositories

import (
	watchlistModel "stock_broker_application/src/models"

	"watchlist/models"

	"gorm.io/gorm"
)

type DeleteWatchListRepository interface {
	DeleteWatchlist(client *gorm.DB, user *models.WatchlistDeleteModel) error
}

type deleteWatchListRepository struct{}

func NewDeleteWatchListRepository() *editWatchListRepository {
	return &editWatchListRepository{}
}

func (repo *editWatchListRepository) DeleteWatchlist(client *gorm.DB, watchlist *models.WatchlistDeleteModel) error {

	err := client.Model(&watchlistModel.Watchlist{}).Where("watchlist_name = ?", watchlist.WatchlistName).Delete(&watchlistModel.Watchlist{}).Error
	if err != nil {
		return err
	}
	return nil
}
