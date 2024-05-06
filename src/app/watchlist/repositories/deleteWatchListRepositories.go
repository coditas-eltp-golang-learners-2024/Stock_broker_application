package repositories

import (
	"log"
	watchlistModel "stock_broker_application/src/models"

	"watchlist/models"

	"gorm.io/gorm"
)

type DeleteWatchListRepository interface {
	DeleteWatchlist(client *gorm.DB, user *models.WatchlistDeleteModel) error
}

type deleteWatchListRepository struct{}

func NewDeleteWatchListRepository() *deleteWatchListRepository {
	return &deleteWatchListRepository{}
}

func (repo *deleteWatchListRepository) DeleteWatchlist(client *gorm.DB, watchlist *models.WatchlistDeleteModel) error {
	log.Println("Before deleting")
	err := client.Model(&watchlistModel.Watchlist{}).Where("watchlist_name = ?", watchlist.WatchlistName).Delete(&watchlistModel.Watchlist{}).Error
	if err != nil {
		return err
	}
	return nil
}
